package bpe

import (
	"fmt"
	"git.missionfocus.com/ours/code/libraries/go/tmetric/client/time_entries"
	"git.missionfocus.com/ours/code/libraries/go/tmetric/models"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/tmetric"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/go-openapi/strfmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

func GetSummedTimeEntries(timeEntries []*models.TimeEntry, isScanner bool) (time.Duration, bool) {
	var totalWorkedHours time.Duration = 0
	projsPassed := true
	for _, e := range timeEntries {
		if !Contains(acceptedProjects, e.ProjectName) && isScanner == true {
			projsPassed = false
		}
		start, err := time.Parse(time.RFC3339, e.StartTime.String())
		if err != nil {
			panic(err)
		}
		end, err := time.Parse(time.RFC3339, e.EndTime.String())
		if err != nil {
			panic(err)
		}
		duration := end.Sub(start)
		totalWorkedHours = totalWorkedHours + duration
	}
	return totalWorkedHours, projsPassed
}

func GetPersonHoursSummary(vaultClient vault.Vault, progress io.Writer, person string) error {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, ' ', 0)

	fmt.Fprintln(progress, "Fetching TMetric Member(s)...")

	scope, _ := tmetric.GetAllTMetricMembers(vaultClient)

	for _, m := range scope.Members {
		emailPointerValue := *m.UserProfile.Email
		profileId := m.UserProfileID
		if person == "" || strings.ToLower(m.UserProfile.UserName) == strings.ToLower(person) || emailPointerValue == person+"@missionfocus.com" {
			calRange := []string{"past day", "week", "month", "lifetime"}
			for _, t := range calRange {
				startDt := strfmt.DateTime(time.Now())
				endDt := strfmt.DateTime(time.Now())
				switch t {
				case "past day":
					startDt = strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Hour * -10))
					endDt = strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Hour * 13).Add(time.Minute * 59))
				case "week":
					startDt = strfmt.DateTime(time.Now().AddDate(0, 0, -5).Add(time.Hour * -10))
					endDt = strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Hour * 13).Add(time.Minute * 59))
				case "month":
					startDt = strfmt.DateTime(time.Now().AddDate(0, -1, 0).Add(time.Hour * -10))
					endDt = strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Hour * 13).Add(time.Minute * 59))
				case "lifetime":
					startDt = strfmt.DateTime(time.Now().AddDate(-100, -1, 0).Add(time.Hour * -10))
					endDt = strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Hour * 13).Add(time.Minute * 59))
				}
				params := time_entries.NewTimeEntriesGetTimeEntriesParams().
					WithAccountID(tmetric.AccountID).
					WithUserProfileID(profileId).
					WithTimeRangeStartTime(&startDt).
					WithTimeRangeEndTime(&endDt)

				timeEntries, _ := tmetric.GetTimeEntriesWithParams(vaultClient, params)
				totalWorkedHours, _ := GetSummedTimeEntries(timeEntries, false)

				fmt.Println(m.UserProfile.UserName+"'s time this", t, "is a total of ", totalWorkedHours)
			}
		} else {
			continue
		}
	}
	return nil
}

func ValidateTMetricTime(vaultClient vault.Vault, progress io.Writer) error {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, ' ', 0)

	scope, _ := tmetric.GetAllTMetricMembers(vaultClient)
	for _, m := range scope.Members {
		profileId := m.UserProfileID
		if m.UserProfile.UserName == "Mission Focus" {
			continue
		}
		fmt.Printf("\n## %s\n\n", m.UserProfile.UserName)

		fmt.Print("Checking username... ")
		if !Contains(acceptedUserNames, m.UserProfile.UserName) {
			fmt.Println("Failed. Unable to find " + m.UserProfile.UserName + " in profile database (username was changed or does not exist).")
		} else {
			fmt.Println("Passed")
		}

		// startDt/startDt test expected be at 10:00 AM each day; making start time = 12:00 AM and end time = 11:59 PM.
		startDt := strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Hour * -10))                    // Start time = 12:00 AM
		endDt := strfmt.DateTime(time.Now().AddDate(0, 0, -1).Add(time.Hour * 13).Add(time.Minute * 59)) // End Time = 11:59 PM
		params := time_entries.NewTimeEntriesGetTimeEntriesParams().
			WithAccountID(tmetric.AccountID).
			WithUserProfileID(profileId).
			WithTimeRangeStartTime(&startDt).
			WithTimeRangeEndTime(&endDt)

		timeEntries, _ := tmetric.GetTimeEntriesWithParams(vaultClient, params)
		totalWorkedHours, projsPassed := GetSummedTimeEntries(timeEntries, true)

		fmt.Print("Checking time projects...")
		if projsPassed {
			fmt.Println("Passed")
		} else {
			fmt.Println(" Failed. User has an invalid project.")
		}

		fmt.Print("Checking total worked hours... ")
		var requiredHours = time.Duration(8) * time.Hour

		if totalWorkedHours < 0 {
			fmt.Println(" Critical Error! User is still logging hours.")
		} else if totalWorkedHours < requiredHours {
			fmt.Print(totalWorkedHours)
			if totalWorkedHours < time.Duration(7)*time.Hour+time.Duration(30)*time.Minute {
				fmt.Println(" Failed. Total hours is less than 7 hours and 30 minutes.")
			} else {
				fmt.Println(" Warning. Total hours are less than 8 hours.")
			}
		} else if totalWorkedHours > time.Duration(8)*time.Hour+time.Duration(30)*time.Minute {
			fmt.Println(" Failed. Total hours is MORE than 8 hours and 30 minutes.")
		} else {
			fmt.Println("Passed")
		}
	}
	return nil
}

//acceptedUserNames this should probably moved to an external file.
var acceptedUserNames = []string{
	"Jacob Stover",
	"Cam Cook",
	"Eric Capito",
	"Wei Zhu",
	"John Kroeker",
	"Matthew Smith",
	"Casey Sault",
	"Alexander Gronowski",
	"Matthew Harbour",
	"Arlo Parker",
	"David Busey",
	"Levi Paulk",
	"Abe Moshekh",
	"Andrew Zaw",
}

var acceptedProjects = []string{"GDAC", "BD-EDM", "PTO", "Overhead"}

// Contains tells whether A [array] contains S [String].
func Contains(a []string, s string) bool {
	for _, n := range a {
		if s == n {
			return true
		}
	}
	return false
}
