# require "formula"

class MfVault < Formula
    desc "Mission Focus distribution of Vault"
    homepage "https://git.missionfocus.com/open-source/mf-vault"
    head "https://git.missionfocus.com/open-source/mf-vault.git"
    # Testing with the README until the actual binary is available
    url "https://git.missionfocus.com/open-source/mf-vault/raw/homebrew-testing/README.md" :nounzip
    sha256 "d1d689727eb4022ed7ccfbb3eb9f085af96183c7c58a57bea5eb1de6199601f6"
    version "0.0.1"
    
    def install
        bin.install "README.md"
    end
end
