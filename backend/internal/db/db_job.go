package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"shift/internal/entity"

	_ "github.com/lib/pq"
)

func (pdb *PostgresDB) CreateJob(job *entity.JobEntity) (*entity.JobEntity, error) {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	jobId, err := pdb.createJob(tx, job)
	if err != nil {
		return nil, err
	}
	res, err := pdb.getJobById(tx, jobId)
	if err != nil {
		logrus.Errorf("getting added job from db: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit job creation in db: %v", err)
		return nil, err
	}
	return res.JobEntity, nil
}

func (pdb *PostgresDB) AssignJobLocation(jobId int, location *entity.JobLocationEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteJobLocation(tx, jobId); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertJobLocation(tx, location); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) AssignJobSkills(jobId int, skills entity.JobSkillsEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteJobSkills(tx, jobId); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertJobSkills(tx, skills); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) AssignJobLanguages(jobId int, locations entity.JobLanguagesEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteJobLanguages(tx, jobId); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertJobLanguages(tx, locations); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) DeleteJob(id int) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	if err := pdb.deleteJobCascade(tx, id); err != nil {
		return err
	}

	query := "update jobs set deleted=true where id=$1"
	res, err := tx.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete job query: %w", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("could not find job %d to delete", id)
	}

	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit company user deletion in db: %v", err)
		return err
	}
	return nil
}

func (pdb *PostgresDB) GetAllJobs() ([]*entity.JobView, error) {
	query := `select jobs.*, loc.city_id, loc.city_name
				from jobs
				left outer join job_locations loc on jobs.id = loc.job_id
				where deleted=false`
	rows, err := pdb.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("fetching all jobs in db: %w", err)
	}
	defer rows.Close()

	res := make([]*entity.JobView, 0)

	for rows.Next() {
		view := new(entity.JobView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan job view from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) GetAllJobsByCompany(companyId int) ([]*entity.JobView, error) {
	query := `select jobs.*, loc.city_id, loc.city_name
				from jobs
				inner join company_users on jobs.creator_id = company_users.company_id
				inner join companies on company_users.company_id = companies.id
				left outer join job_locations loc on jobs.id = loc.job_id
				where companies.id = $1 and jobs.deleted=false`
	rows, err := pdb.db.Queryx(query, companyId)
	if err != nil {
		return nil, fmt.Errorf("fetching all jobs for company id %d in db: %w", companyId, err)
	}
	defer rows.Close()

	res := make([]*entity.JobView, 0)

	for rows.Next() {
		view := new(entity.JobView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan job view from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) GetJobById(id int) (*entity.JobView, error) {
	return pdb.getJobById(pdb.db, id)
}

func (pdb *PostgresDB) GetSkillsByJobId(jobId int) (entity.JobSkillsEntity, error) {
	query := `select job_skills.*
				from job_skills
				inner join jobs on job_skills.job_id = jobs.id
				where jobs.id = $1`
	rows, err := pdb.db.Queryx(query, jobId)
	if err != nil {
		return nil, fmt.Errorf("fetching job %d skils in db: %w", jobId, err)
	}
	defer rows.Close()

	res := make(entity.JobSkillsEntity, 0)

	for rows.Next() {
		view := new(entity.JobSkillEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan job skill from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}
func (pdb *PostgresDB) GetLanguagesByJobId(jobId int) (entity.JobLanguagesEntity, error) {
	query := `select job_languages.*
				from job_languages
				inner join jobs on job_languages.job_id = jobs.id
				where jobs.id = $1`
	rows, err := pdb.db.Queryx(query, jobId)
	if err != nil {
		return nil, fmt.Errorf("fetching job %d languages in db: %w", jobId, err)
	}
	defer rows.Close()

	res := make(entity.JobLanguagesEntity, 0)

	for rows.Next() {
		view := new(entity.JobLanguageEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan job language from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) getJobById(tx sqlx.Queryer, id int) (*entity.JobView, error) {
	query := `select jobs.*, loc.city_id, loc.city_name
				from jobs
				left outer join job_locations loc on jobs.id = loc.job_id
				where jobs.id = $1 and deleted=false`
	rows, err := tx.Queryx(query, id)
	if err != nil {
		return nil, fmt.Errorf("fetching job id=%d in db: %w", id, err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.JobView)
		if err := rows.StructScan(view); err != nil {
			return nil, fmt.Errorf("failed to scan job view from db row: %v", err)
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find job: id=%d", id)
}

func (pdb *PostgresDB) createJob(tx NamedQuerier, job *entity.JobEntity) (int, error) {
	query := `insert into jobs
				(
				 	title,
					creator_id,
					experience_from,
					experience_to,
					job_type,
					employment_level_from,
					employment_level_to,
					overview,
					role_responsibilities,
					candidate_description,
					location_type,
					salary_range_from,
					salary_range_to,
					benefits,
					start_date
				)
				values (
				 	:title,
					:creator_id,
					:experience_from,
					:experience_to,
					:job_type,
					:employment_level_from,
					:employment_level_to,
					:overview,
					:role_responsibilities,
					:candidate_description,
					:location_type,
					:salary_range_from,
					:salary_range_to,
					:benefits,
					:start_date
				)
				returning id`
	jobId, err := PreparedQuery(tx, query, job)
	if err != nil {
		logrus.Debugf("failed to insert job in db: %v", err)
		return 0, err
	}
	return jobId, nil
}

func (pdb *PostgresDB) insertJobLocation(tx NamedQuerier, record *entity.JobLocationEntity) error {
	query := `insert into job_locations (job_id, city_id, city_name) values (:job_id, :city_id, :city_name)`
	if _, err := tx.NamedExec(query, record); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteJobLocation(tx sqlx.Execer, jobId int) error {
	query := `delete from job_locations where job_id = $1`
	if _, err := tx.Exec(query, jobId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) insertJobSkills(tx NamedQuerier, records entity.JobSkillsEntity) error {
	query := `insert into job_skills (job_id, name) values (:job_id, :name)`
	for _, record := range records {
		if _, err := tx.NamedExec(query, record); err != nil {
			return err
		}
	}
	return nil
}

func (pdb *PostgresDB) deleteJobSkills(tx sqlx.Execer, jobId int) error {
	query := `delete from job_skills where job_id = $1`
	if _, err := tx.Exec(query, jobId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) insertJobLanguages(tx NamedQuerier, records entity.JobLanguagesEntity) error {
	query := `insert into job_languages (job_id, language_id, language_name, language_short_name) values (:job_id, :language_id, :language_name, :language_short_name)`
	for _, record := range records {
		if _, err := tx.NamedExec(query, record); err != nil {
			return err
		}
	}
	return nil
}

func (pdb *PostgresDB) deleteJobLanguages(tx sqlx.Execer, jobId int) error {
	query := `delete from job_languages where job_id = $1`
	if _, err := tx.Exec(query, jobId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteJobsByCreatorId(tx QueryerExecer, creatorId int) error {
	query := "update jobs set deleted=true where id=$1 returning id"
	rows, err := tx.Queryx(query, creatorId)
	if err != nil {
		return fmt.Errorf("delete job by creator id query: %w", err)
	}
	defer rows.Close()

	ids := make([]int, 0)
	for rows.Next() {
		view := new(struct {
			Id int `db:"id"`
		})
		if err := rows.StructScan(&view); err != nil {
			logrus.Errorf("failed to scan job id from db row: %v", err)
			return err
		}
		ids = append(ids, view.Id)
	}

	for _, id := range ids {
		if err := pdb.deleteJobCascade(tx, id); err != nil {
			return err
		}
	}

	return nil
}

func (pdb *PostgresDB) deleteJobCascade(tx sqlx.Execer, id int) error {
	if err := pdb.deleteJobLocation(tx, id); err != nil {
		return fmt.Errorf("delete job location: %w", err)
	}

	if err := pdb.deleteJobSkills(tx, id); err != nil {
		return fmt.Errorf("delete job skills: %w", err)
	}

	if err := pdb.deleteJobLanguages(tx, id); err != nil {
		return fmt.Errorf("delete job languages: %w", err)
	}
	return nil
}
