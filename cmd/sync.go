package cmd

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/OpsVerseIO/argocd-actions/internal/argocd"
	ctrl "github.com/OpsVerseIO/argocd-actions/internal/controller"
)

// Sync syncs given application.
func Sync() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Sync ArgoCD application.",
		RunE: func(cmd *cobra.Command, args []string) error {
			address, _ := cmd.Flags().GetString("address")
			token, _ := cmd.Flags().GetString("token")
			application, _ := cmd.Flags().GetString("application")
			disableTlsVerification, _ := cmd.Flags().GetString("disableTlsVerification")
			disableTlsVerificationBool, _ := strconv.ParseBool(disableTlsVerification)

			api := argocd.NewAPI(&argocd.APIOptions{
				Address:                address,
				Token:                  token,
				DisableTlsVerification: disableTlsVerificationBool,
			})

			controller := ctrl.NewController(api)

			err := controller.Sync(application)
			if err != nil {
				return err
			}

			log.Infof("Application %s synced", application)

			return nil
		},
	}

	cmd.Flags().String("application", "", "ArgoCD application name")

	if err := cmd.MarkFlagRequired("application"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	return cmd
}
