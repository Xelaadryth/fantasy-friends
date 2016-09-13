package database

// prepareUserStatements readies SQL queries for later use
import "github.com/jackc/pgx"

//FantasyTeam contains the summoner IDs and meta information for a fantasy team
type FantasyTeam struct {
	ID       int64
	Owner    int64
	Name     string
	Position int
	Top      int64
	Jungle   int64
	Mid      int64
	Bottom   int64
	Support  int64
}

//Queries that are prepared for easy calling
const (
	QueryGetTeams   = "getTeams"
	QueryInsertTeam = "insertTeam"
	QueryUpdateTeam = "updateTeam"
	QueryDeleteTeam = "deleteTeam"
)

//GetTeams for the given user id from the database if they exist
func GetTeams(userID int64) (*[]FantasyTeam, error) {
	//TODO: Update this if we expect more than 1 team per player in the future
	teams := make([]FantasyTeam, 0, 1)
	rows, _ := DBConnectionPool.Query(QueryGetTeams, userID)

	for rows.Next() {
		team := FantasyTeam{}
		rows.Scan(
			&(team.ID),
			&(team.Owner),
			&(team.Name),
			&(team.Position),
			&(team.Top),
			&(team.Jungle),
			&(team.Mid),
			&(team.Bottom),
			&(team.Support),
		)

		teams = append(teams, team)
	}

	return &teams, rows.Err()
}

//AddTeam to database, errors if team already exists in that slot
func AddTeam(owner int64, teamName string, position int,
	top int64, jungle int64, mid int64, bottom int64, support int64) error {
	//Get the cache IDs
	_, err := DBConnectionPool.Exec(
		QueryInsertTeam,
		owner,
		teamName,
		position,
		top,
		jungle,
		mid,
		bottom,
		support,
	)

	return err
}

//UpdateTeam that already exists in the database
func UpdateTeam(teamID int64, owner int64, teamName string, position int,
	top int64, jungle int64, mid int64, bottom int64, support int64) error {
	//Get the cache IDs
	_, err := DBConnectionPool.Exec(
		QueryUpdateTeam,
		teamID,
		owner,
		teamName,
		position,
		top,
		jungle,
		mid,
		bottom,
		support,
	)

	return err
}

//DeleteTeam that already exists in the database
func DeleteTeam(teamID int64) error {
	_, err := DBConnectionPool.Exec(
		QueryDeleteTeam,
		teamID,
	)

	return err
}

func prepareTeamStatements(conn *pgx.Conn) error {
	_, err := conn.Prepare(QueryGetTeams, `
		SELECT id, owner, team_name, position, top, jungle, mid, bottom, support
		FROM fantasy_friends.fantasy_team
		WHERE owner=$1
  `)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(QueryInsertTeam, `
		INSERT INTO fantasy_friends.fantasy_team (owner, team_name, position, top, jungle, mid, bottom, support)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`)
	if err != nil {
		return err
	}

	_, err = conn.Prepare(QueryUpdateTeam, `
		UPDATE fantasy_friends.fantasy_team
		SET (owner, team_name, position, top, jungle, mid, bottom, support) = ($2, $3, $4, $5, $6, $7, $8, $9)
		WHERE id = $1
	`)

	_, err = conn.Prepare(QueryDeleteTeam, `
		DELETE FROM fantasy_friends.fantasy_team
		WHERE id = $1
	`)

	return err
}