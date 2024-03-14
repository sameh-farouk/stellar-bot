module github.com/sameh-farouk/stellar-bot

go 1.19

replace github.com/sameh-farouk/stellar-bot/config/ => /home/sameh/projects/stellar-bot/internal/config/config.go

replace github.com/sameh-farouk/stellar-bot/monitoring/ => /home/sameh/projects/stellar-bot/internal/monitoring/monitoring.go

replace github.com/sameh-farouk/stellar-bot/stellar/ => /home/sameh/projects/stellar-bot/internal/stellar/stellar.go

replace github.com/sameh-farouk/stellar-bot/trading/ => /home/sameh/projects/stellar-bot/internal/trading/trading.go

replace github.com/sameh-farouk/stellar-bot/user/ => /home/sameh/projects/stellar-bot/internal/user/user.go

require (
	github.com/spf13/viper v1.18.2
	github.com/stellar/go v0.0.0-20240311093440-6fe5270b6f0b
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-chi/chi v4.1.2+incompatible // indirect
	github.com/go-errors/errors v1.5.1 // indirect
	github.com/gorilla/schema v1.2.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/manucorporat/sse v0.0.0-20160126180136-ee05b128a739 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/segmentio/go-loggly v0.5.1-0.20171222203950-eb91657e62b2 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stellar/go-xdr v0.0.0-20231122183749-b53fb00bcac2 // indirect
	github.com/stretchr/objx v0.5.1 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
