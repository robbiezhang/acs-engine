package cmd

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var _ = Describe("the upgrade command", func() {

	It("should create an upgrade command", func() {
		output := newUpgradeCmd()

		Expect(output.Use).Should(Equal(upgradeName))
		Expect(output.Short).Should(Equal(upgradeShortDescription))
		Expect(output.Long).Should(Equal(upgradeLongDescription))
		Expect(output.Flags().Lookup("location")).NotTo(BeNil())
		Expect(output.Flags().Lookup("resource-group")).NotTo(BeNil())
		Expect(output.Flags().Lookup("deployment-dir")).NotTo(BeNil())
		Expect(output.Flags().Lookup("upgrade-version")).NotTo(BeNil())
	})

	It("should validate an upgrade command", func() {
		r := &cobra.Command{}

		cases := []struct {
			uc          *upgradeCmd
			expectedErr error
		}{
			{
				uc: &upgradeCmd{
					resourceGroupName:   "",
					deploymentDirectory: "_output/test",
					upgradeVersion:      "1.8.9",
					location:            "centralus",
					timeoutInMinutes:    60,
					authArgs: authArgs{
						rawSubscriptionID: "99999999-0000-0000-0000-000000000000",
					},
				},
				expectedErr: fmt.Errorf("--resource-group must be specified"),
			},
			{
				uc: &upgradeCmd{
					resourceGroupName:   "test",
					deploymentDirectory: "_output/test",
					upgradeVersion:      "1.8.9",
					location:            "",
					timeoutInMinutes:    60,
					authArgs: authArgs{
						rawSubscriptionID: "99999999-0000-0000-0000-000000000000",
					},
				},
				expectedErr: fmt.Errorf("--location must be specified"),
			},
			{
				uc: &upgradeCmd{
					resourceGroupName:   "test",
					deploymentDirectory: "_output/test",
					upgradeVersion:      "",
					location:            "southcentralus",
					timeoutInMinutes:    60,
					authArgs: authArgs{
						rawSubscriptionID: "99999999-0000-0000-0000-000000000000",
					},
				},
				expectedErr: fmt.Errorf("--upgrade-version must be specified"),
			},
			{
				uc: &upgradeCmd{
					resourceGroupName:   "test",
					deploymentDirectory: "",
					upgradeVersion:      "1.9.0",
					location:            "southcentralus",
					timeoutInMinutes:    60,
					authArgs: authArgs{
						rawSubscriptionID: "99999999-0000-0000-0000-000000000000",
					},
				},
				expectedErr: fmt.Errorf("--deployment-dir must be specified"),
			},
			{
				uc: &upgradeCmd{
					resourceGroupName:   "test",
					deploymentDirectory: "",
					upgradeVersion:      "1.9.0",
					location:            "southcentralus",
					timeoutInMinutes:    60,
					authArgs: authArgs{
						rawSubscriptionID: "99999999-0000-0000-0000-000000000000",
					},
				},
				expectedErr: fmt.Errorf("--deployment-dir must be specified"),
			},
			{
				uc: &upgradeCmd{
					resourceGroupName:   "test",
					deploymentDirectory: "_output/mydir",
					upgradeVersion:      "1.9.0",
					location:            "southcentralus",
					authArgs:            authArgs{},
				},
				expectedErr: fmt.Errorf("--subscription-id is required (and must be a valid UUID)"),
			},
			{
				uc: &upgradeCmd{
					resourceGroupName:   "test",
					deploymentDirectory: "_output/mydir",
					upgradeVersion:      "1.9.0",
					location:            "southcentralus",
					authArgs: authArgs{
						rawSubscriptionID:   "99999999-0000-0000-0000-000000000000",
						RawAzureEnvironment: "AzurePublicCloud",
						AuthMethod:          "device",
					},
				},
				expectedErr: nil,
			},
		}

		for _, c := range cases {
			err := c.uc.validate(r)
			if c.expectedErr != nil && err != nil {
				Expect(err.Error()).To(Equal(c.expectedErr.Error()))
			} else {
				Expect(err).To(BeNil())
				Expect(c.expectedErr).To(BeNil())
			}
		}

	})

})
