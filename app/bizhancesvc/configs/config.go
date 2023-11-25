package configs

import "github.com/mkideal/cli"

const AppVersion = "v0.0.1"

type ServerConfig struct {
	cli.Helper
	ServerPort         int    `cli:"*port" usage:"Application database URI" dft:"$BIZHANCESVC_PORT"`
	DBURI              string `cli:"*dburi" usage:"Application database URI" dft:"$BIZHANCESVC_DBURI"`
	JWTSecret          string `cli:"*jwtsecret" usage:"Secret JWT token" dft:"$BIZHANCESVC_JWT_SECRET"`
	SentryDSN          string `cli:"*sentrydsn" usage:"Sentry DSN" dft:"$BIZHANCESVC_SENTRY_DSN"`
	SMTPHostURL        string `cli:"*smtphosturl" usage:"" dft:"$BIZHANCESVC_SMTP_HOST_URL"`
	SMTPHostPORT       string `cli:"*smtphostport" usage:"" dft:"$BIZHANCESVC_SMTP_HOST_PORT"`
	SMTPUsername       string `cli:"*smtpusername" usage:"" dft:"$BIZHANCESVC_SMTP_USERNAME"`
	SMTPPassword       string `cli:"*smtppassword" usage:"" dft:"$BIZHANCESVC_SMTP_PASSWORD"`
	EmailDefaultSender string `cli:"*emaildefaultsender" usage:"" dft:"$BIZHANCESVC_EMAIL_SENDER"`
}
