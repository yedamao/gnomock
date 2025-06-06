package mariadb

// Option is an optional configuration of this Gnomock preset. Use available
// Options to configure the container
type Option func(*P)

// WithUser creates a new superuser with the provided credentials in the
// container. If not used, the default credentials are gnomock:gnoria
func WithUser(user, password string) Option {
	return func(p *P) {
		p.User = user
		p.Password = password
	}
}

// WithDatabase creates a database with the provided name in the container. If
// not provided, "mydb" is used by default.  WithQueries, if provided, runs
// against the new database
func WithDatabase(db string) Option {
	return func(p *P) {
		p.DB = db
	}
}

// WithQueries executes the provided queries against the database created with
// WithDatabase, or against default "mydb" database
func WithQueries(queries ...string) Option {
	return func(p *P) {
		p.Queries = append(p.Queries, queries...)
	}
}

// WithQueriesFile sets a file name to read initial queries from. Queries from
// this file are executed before any other queries provided in WithQueries
func WithQueriesFile(file string) Option {
	return func(p *P) {
		p.QueriesFiles = append(p.QueriesFiles, file)
	}
}

// WithImageName sets image name.
func WithImageName(imageName string) Option {
	return func(p *P) {
		p.ImageName = imageName
	}
}

// WithVersion sets image version.
func WithVersion(version string) Option {
	return func(o *P) {
		o.Version = version
	}
}
