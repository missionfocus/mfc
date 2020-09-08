package bpe

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"git.missionfocus.com/ours/code/libraries/go/tmetric/client/time_entries"
	"git.missionfocus.com/ours/code/libraries/go/tmetric/models"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/tmetric"
	"git.missionfocus.com/ours/code/tools/mfc/pkg/vault"
	"github.com/go-openapi/strfmt"
)

var year, month, day = time.Now().Date()
var today = time.Date(year, month, day, 0, 0, 0, 0, time.Local)
var yesterday = today.AddDate(0, 0, -1)
var monday = today.AddDate(0, 0, -1*(int(today.Weekday())-1))
var beginOfMonth = time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
var lifeTime = today.AddDate(-100, 0, 0)
var startDt = strfmt.DateTime(yesterday)
var endDt = strfmt.DateTime(today)

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
	members, _ := tmetric.GetAllTMetricMembers(vaultClient)

	for _, m := range members {
		emailPointerValue := *m.UserProfile.Email
		profileId := m.UserProfileID
		if person == "" || strings.ToLower(m.UserProfile.UserName) == strings.ToLower(person) || emailPointerValue == person+"@missionfocus.com" {
			calRange := []string{"past day", "week", "month", "lifetime"}
			for _, t := range calRange {
				switch t {
				case "past day":
					startDt = strfmt.DateTime(yesterday)
				case "week":
					startDt = strfmt.DateTime(monday)
				case "month":
					startDt = strfmt.DateTime(beginOfMonth)
				case "lifetime":
					startDt = strfmt.DateTime(lifeTime)
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
			fmt.Println()
			continue
		}
	}
	return nil
}

func ValidateTMetricTime(vaultClient vault.Vault) error {
	members, _ := tmetric.GetAllTMetricMembers(vaultClient)
	for _, m := range members {
		profileId := m.UserProfileID
		if m.UserProfile.UserName == "Mission Focus" {
			continue
		}
		minimumHours := time.Duration(7)*time.Hour + time.Duration(30)*time.Minute
		maximumHours := time.Duration(8)*time.Hour + time.Duration(30)*time.Minute
		params := time_entries.NewTimeEntriesGetTimeEntriesParams().
			WithAccountID(tmetric.AccountID).
			WithUserProfileID(profileId).
			WithTimeRangeStartTime(&startDt).
			WithTimeRangeEndTime(&endDt)
		timeEntries, _ := tmetric.GetTimeEntriesWithParams(vaultClient, params)
		totalWorkedHours, projsPassed := GetSummedTimeEntries(timeEntries, true)

		fmt.Printf("\n## %s\n", m.UserProfile.UserName)
		fmt.Print("Checking username... ")
		if !Contains(acceptedUserNames, m.UserProfile.UserName) {
			fmt.Println("Failed. Unable to find " + m.UserProfile.UserName + " in profile database (username was changed or does not exist).")
		} else {
			fmt.Println("Passed")
		}
		fmt.Print("Checking time projects... ")
		if projsPassed {
			fmt.Println("Passed")
		} else {
			fmt.Println("Failed. User has an invalid project.")
		}
		fmt.Print("Checking total worked hours... ")
		if totalWorkedHours < 0 {
			fmt.Println("Error: User is still logging hours.")
		}
		fmt.Print(totalWorkedHours, " ")
		if totalWorkedHours < minimumHours {
			fmt.Println("Failed. Total hours is less than 7 hours and 30 minutes.")
		} else if totalWorkedHours > maximumHours {
			fmt.Println("Failed. Total hours is more than 8 hours amd 30 minutes.")
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

var acceptedProjects = []string{"GDAC", "BD", "BD-EDM", "Holiday", "PTO", "IRAD", "Overhead"}

// Contains tells whether A [array] contains S [String].
func Contains(a []string, s string) bool {
	for _, n := range a {
		if s == n {
			return true
		}
	}
	return false
}
