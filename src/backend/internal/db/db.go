package db

import (
	"fmt"
	"log"
	"os"
	"shift/internal/entity"

	"github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sqlx.DB
}

func NewPostgresDB() *PostgresDB {
	db, err := sqlx.Connect("postgres", os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	return &PostgresDB{
		db: db,
	}
}

func (pdb *PostgresDB) DeleteAdminUser(id int) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	if err := pdb.deleteJobsByCreatorId(tx, id); err != nil {
		return fmt.Errorf("delete company user jobs: %w", err)
	}

	query := "update users set state='DELETED' where id=$1 and kind='ADMIN'"
	res, err := pdb.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete user of admin kind: %w", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("could not find admin %d to delete", id)
	}

	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit admin deletion in db: %v", err)
		return err
	}
	return nil
}

func (pdb *PostgresDB) DeleteAssociationUser(id int) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	query := "delete from association_users where user_id=$1"
	res, err := tx.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete association user: %w", err)
	}

	query = "update users set state='DELETED' where id=$1 and kind='ASSOCIATION'"
	res, err = tx.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete user of association user kind: %w", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("could not find association user %d to delete", id)
	}

	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit association user deletion in db: %v", err)
		return err
	}
	return nil
}

func (pdb *PostgresDB) DeleteCandidateUser(id int) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	if err := pdb.deleteCandidateSkillsByUserId(tx, id); err != nil {
		return fmt.Errorf("deleting candidate skills: % by user idv", err)
	}

	if err := pdb.deleteCandidateSpokenLanguagesByUserId(tx, id); err != nil {
		return fmt.Errorf("deleting candidate spoken languages:  by user id%v", err)
	}

	if err := pdb.deleteCandidateSeekLocationsByUserId(tx, id); err != nil {
		return fmt.Errorf("deleting candidate seek locations:  by user id%v", err)
	}

	if err := pdb.deleteCandidateEducationHistoryListByUserId(tx, id); err != nil {
		return fmt.Errorf("deleting candidate education history list by user id: %v", err)
	}

	if err := pdb.deleteCandidateEmploymentHistoryListByUserId(tx, id); err != nil {
		return fmt.Errorf("deleting candidate employment history list by user id: %v", err)
	}

	query := "delete from candidates where user_id=$1"
	res, err := tx.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete candidate: %w", err)
	}

	query = "update users set state='DELETED' where id=$1 and kind='CANDIDATE'"
	res, err = tx.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete user of candidate kind: %w", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("could not find candidate %d to delete", id)
	}

	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit candidate deletion in db: %v", err)
		return err
	}
	return nil
}

func (pdb *PostgresDB) DeleteCompanyUser(id int) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	if err := pdb.deleteJobsByCreatorId(tx, id); err != nil {
		return fmt.Errorf("delete company user jobs: %w", err)
	}

	query := "delete from company_users where user_id=$1"
	res, err := tx.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete company user: %w", err)
	}

	query = "update users set state='DELETED' where id=$1 and kind='COMPANY'"
	res, err = tx.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete user of company user kind: %w", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("could not find company user %d to delete", id)
	}

	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit company user deletion in db: %v", err)
		return err
	}
	return nil
}

func (pdb *PostgresDB) GetUserRecord(id int) (*entity.UserRecordView, error) {
	query := `select id, kind, email, state, created_at
				from users
				where id = $1 and state = 'ACTIVE'`
	rows, err := pdb.db.Queryx(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.UserRecordView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user record view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find user record view: id=%d", id)
}

func (pdb *PostgresDB) GetUserRecordByCompanyUserId(companyUserId int) (*entity.UserRecordView, error) {
	query := `select users.id, kind, email, state, created_at
				from users
				inner join company_users on users.id = company_users.user_id
				where company_users.id = $1 and state = 'ACTIVE'`
	rows, err := pdb.db.Queryx(query, companyUserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.UserRecordView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user record view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find user record view: company_user_id=%d", companyUserId)
}

func (pdb *PostgresDB) GetUserRecordsByCompanyId(companyId int) ([]*entity.UserRecordView, error) {
	query := `select users.id, kind, email, state, users.created_at
				from users
				inner join company_users on users.id = company_users.user_id
				inner join companies on company_users.company_id = companies.id
				where companies.id = $1 and state = 'ACTIVE'`
	rows, err := pdb.db.Queryx(query, companyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*entity.UserRecordView, 0)

	for rows.Next() {
		view := new(entity.UserRecordView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user record view from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) GetUserRecordsByAssociationId(associationId int) ([]*entity.UserRecordView, error) {
	query := `select users.id, kind, email, state, users.created_at
				from users
				inner join association_users on users.id = association_users.user_id
				inner join associations on association_users.association_id = associations.id
				where associations.id = $1 and state = 'ACTIVE'`
	rows, err := pdb.db.Queryx(query, associationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*entity.UserRecordView, 0)

	for rows.Next() {
		view := new(entity.UserRecordView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user record view from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) GetUserRecordByEmail(email string) (*entity.UserRecordView, error) {
	query := `select id, kind, email, state, created_at
				from users
				where email = $1 and state = 'ACTIVE'`
	rows, err := pdb.db.Queryx(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.UserRecordView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user record view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find user record view: email=%s", email)
}

func (s *PostgresDB) GetAssociationRecord(id int) (*entity.AssociationEntity, error) {
	query := `select * from associations where id = $1 and deleted = false`
	rows, err := s.db.Queryx(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.AssociationEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan association record view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find user record view: id=%v", id)
}

func (pdb *PostgresDB) GetProfileByEmail(email string) (*entity.UserProfileView, error) {
	query := `select
    				users.id,
    				kind,
    				first_name,
    				last_name,
    				preferred_name,
    				email,
    				state,
    				created_at,
    				image_url,
    				association_users.association_id,
    				company_users.company_id
				from users
				left outer join user_photos on users.id = user_photos.user_id
				left outer join association_users on users.id = association_users.user_id
				left outer join company_users on users.id = company_users.user_id
				where email = $1`
	rows, err := pdb.db.Queryx(query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.UserProfileView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user record view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find user record view: email=%s", email)
}

func (pdb *PostgresDB) GetAllUsers() ([]*entity.UserItemView, error) {
	res := make([]*entity.UserItemView, 0)

	query := `select
    				users.id,
					users.kind,
					users.first_name,
					users.last_name,
					users.preferred_name,
					users.email,
					users.phone_number,
					users.birth_date,
					users.linkedin_url,
					users.github_url,
					users.portfolio_url,
					users.state,
					users.created_at,
    				association_users.id as association_user_id,
    				association_users.association_id,
    				association_users.role as association_role,
    				candidates.id as candidate_id,
    				candidates.years_of_experience,
    				candidates.job_status,
    				candidates.seek_job_type,
    				candidates.seek_company_size,
    				candidates.seek_location_type,
    				candidates.seek_salary,
    				candidates.seek_values,
    				candidates.work_permit,
    				candidates.notice_period,
    				company_users.id as company_user_id,
    				company_users.company_id,
    				company_users.role as company_role,
    				user_photos.image_url,
					candidate_cvs.cv_url,
					candidate_videos.video_url
				from users
				left outer join candidates on users.id = candidates.user_id
				left outer join association_users on users.id = association_users.user_id
				left outer join company_users on users.id = company_users.user_id
				left outer join user_photos on users.id = user_photos.user_id
				left outer join candidate_cvs on candidates.id = candidate_cvs.candidate_id
				left outer join candidate_videos on candidates.id = candidate_videos.candidate_id
    `
	rows, err := pdb.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("fetching users in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.UserItemView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user view from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) GetAllAssociations() ([]*entity.AssociationEntity, error) {
	res := make([]*entity.AssociationEntity, 0)

	query := `select * from associations where deleted = false`

	rows, err := pdb.db.Queryx(query)

	if err != nil {
		return nil, fmt.Errorf("fetching associations in db: %w", err)
	}

	for rows.Next() {
		view := new(entity.AssociationEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Debugf("failed to scan association view from db record: %v", err)
			return nil, err
		}
		res = append(res, view)
	}
	fmt.Println(res)
	return res, nil
}

func (pdb *PostgresDB) GetAllInvitations() ([]*entity.InvitationEntity, error) {
	res := make([]*entity.InvitationEntity, 0)

	query := `select * from invites`
	rows, err := pdb.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("fetching invitations in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.InvitationEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user view from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) GetInvitationByEmail(email string) (*entity.InvitationEntity, error) {
	query := `select * from invites where email=$1`
	rows, err := pdb.db.Queryx(query, email)
	if err != nil {
		return nil, fmt.Errorf("fetching invitation by email in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.InvitationEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find invitation: email=%s", email)
}

func (pdb *PostgresDB) GetUserById(id int) (*entity.UserItemView, error) {
	query := `select
    				users.id,
					users.kind,
					users.first_name,
					users.last_name,
					users.preferred_name,
					users.email,
					users.phone_number,
					users.birth_date,
					users.linkedin_url,
					users.github_url,
					users.portfolio_url,
					users.state,
					users.created_at,
    				user_photos.image_url
				from users
				left outer join user_photos on users.id = user_photos.user_id
				where users.id = $1
    `
	rows, err := pdb.db.Queryx(query, id)
	if err != nil {
		return nil, fmt.Errorf("fetching user id=%d in db: %w", id, err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.UserItemView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find user: id=%d", id)
}

func (pdb *PostgresDB) GetAssociationById(id int) (*entity.AssociationEntity, error) {
	query := `select * from associations where id = $1 and deleted = false`
	rows, err := pdb.db.Queryx(query, id)
	if err != nil {
		return nil, fmt.Errorf("fetching association id=%d in db: %w", id, err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.AssociationEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find user: id=%d", id)
}

func (pdb *PostgresDB) GetAssociationUserByUserId(id int) (*entity.UserItemView, error) {
	query := `select
    				users.id,
					users.kind,
					users.first_name,
					users.last_name,
					users.preferred_name,
					users.email,
					users.phone_number,
					users.birth_date,
					users.linkedin_url,
					users.github_url,
					users.portfolio_url,
					users.state,
					users.created_at,
    				association_users.id as association_user_id,
    				association_users.association_id,
    				association_users.role as association_role,
    				user_photos.image_url
				from users
				inner join association_users on users.id = association_users.user_id
				left outer join user_photos on users.id = user_photos.user_id
				where users.id = $1
    `
	rows, err := pdb.db.Queryx(query, id)
	if err != nil {
		return nil, fmt.Errorf("fetching association user id=%d in db: %w", id, err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.UserItemView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find association user: id=%d", id)
}

func (pdb *PostgresDB) GetCandidateByUserId(id int) (*entity.UserItemView, error) {
	query := `select
    				users.id,
					users.kind,
					users.first_name,
					users.last_name,
					users.preferred_name,
					users.email,
					users.phone_number,
					users.birth_date,
					users.linkedin_url,
					users.github_url,
					users.portfolio_url,
					users.state,
					users.created_at,
    				candidates.id as candidate_id,
    				candidates.years_of_experience,
    				candidates.job_status,
    				candidates.seek_job_type,
    				candidates.seek_company_size,
    				candidates.seek_location_type,
    				candidates.seek_salary,
    				candidates.seek_values,
    				candidates.work_permit,
    				candidates.notice_period,
    				user_photos.image_url,
					candidate_cvs.cv_url,
					candidate_videos.video_url
				from users
				left outer join candidates on users.id = candidates.user_id
				left outer join user_photos on users.id = user_photos.user_id
				left outer join candidate_cvs on candidates.id = candidate_cvs.candidate_id
				left outer join candidate_videos on candidates.id = candidate_videos.candidate_id
				where users.id = $1
    `
	rows, err := pdb.db.Queryx(query, id)
	if err != nil {
		return nil, fmt.Errorf("fetching candidate id=%d in db: %w", id, err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.UserItemView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find candidate: id=%d", id)
}

func (pdb *PostgresDB) GetCompanyUserByUserId(id int) (*entity.UserItemView, error) {
	query := `select
    				users.id,
					users.kind,
					users.first_name,
					users.last_name,
					users.preferred_name,
					users.email,
					users.phone_number,
					users.birth_date,
					users.linkedin_url,
					users.github_url,
					users.portfolio_url,
					users.state,
					users.created_at,
    				company_users.id as company_user_id,
    				company_users.company_id,
    				company_users.role as company_role,
    				user_photos.image_url
				from users
				left outer join company_users on users.id = company_users.user_id
				left outer join user_photos on users.id = user_photos.user_id
				where users.id = $1
    `
	rows, err := pdb.db.Queryx(query, id)
	if err != nil {
		return nil, fmt.Errorf("fetching company user id=%d in db: %w", id, err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.UserItemView)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan user view from db row: %v", err)
			return nil, err
		}
		return view, nil
	}

	return nil, fmt.Errorf("could not find company user: id=%d", id)
}

func (pdb *PostgresDB) CreateUser(user *entity.UserEntity) (*entity.UserEntity, error) {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	userId, err := pdb.createUser(tx, user)
	if err != nil {
		return nil, err
	}
	user, err = pdb.getUserById(tx, userId)
	if err != nil {
		logrus.Errorf("getting added user from db: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit user creation in db: %v", err)
		return nil, err
	}
	return user, nil
}

func (pdb *PostgresDB) EditUser(id int, user *entity.UserEntity) (*entity.UserEntity, error) {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	userId, err := pdb.editUser(tx, id, user)
	if err != nil {
		return nil, err
	}
	res, err := pdb.getUserById(tx, userId)
	if err != nil {
		logrus.Errorf("getting edited user from db: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit user update in db: %v", err)
		return nil, err
	}
	return res, nil
}

func (pdb *PostgresDB) CreateAssociation(assoc *entity.AssociationEntity) (*entity.AssociationEntity, error) {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	assocId, err := pdb.createAssociation(tx, assoc)
	if err != nil {
		return nil, err
	}
	assoc, err = pdb.getAssociationById(tx, assocId)
	if err != nil {
		logrus.Errorf("getting added associations from db: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit associations creation in db: %v", err)
		return nil, err
	}
	return assoc, nil
}

func (pdb *PostgresDB) AssignAssociationLogo(id int, logoUrl string) error {
	query := `update associations set logo_url=$1 where id=$2`

	_, err := pdb.db.Queryx(query, logoUrl, id)
	if err != nil {
		return fmt.Errorf("could not assign logo %s to association %d: %w", logoUrl, id, err)
	}
	return nil
}

func (pdb *PostgresDB) DeleteAssociation(id int) error {
	query := "update associations set deleted=true WHERE id = $1"
	res, err := pdb.db.Exec(query, id)

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("association not found")
	}
	return nil
}

func (pdb *PostgresDB) CreateInvitation(inv *entity.InvitationEntity) (*entity.InvitationEntity, error) {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	invId, err := pdb.createInvitation(tx, inv)
	if err != nil {
		return nil, err
	}
	inv, err = pdb.getInvitationById(tx, invId)
	if err != nil {
		logrus.Errorf("getting added invitations from db: %v", err)
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit invitations creation in db: %v", err)
		return nil, err
	}

	return inv, nil
}

func (pdb *PostgresDB) UpdateInvitationState(id int, state string) error {
	query := `update invites set state=$2 where id=$1`
	if _, err := pdb.db.Queryx(query, id, state); err != nil {
		return fmt.Errorf("executing query: %w", err)
	}
	return nil
}

func (pdb *PostgresDB) SetInvitationTicket(id int, ticket string) error {
	query := `update invites set state='PENDING', ticket=$2 where id=$1`
	if _, err := pdb.db.Queryx(query, id, ticket); err != nil {
		return fmt.Errorf("executing query: %w", err)
	}
	return nil
}

func (pdb *PostgresDB) getAssociationById(tx sqlx.Queryer, id int) (*entity.AssociationEntity, error) {
	query := `select * from associations where id = $1 and deleted = false`
	rows, err := tx.Queryx(query, id)
	if err != nil {
		logrus.Debugf("failed to get association with id=%d in db: %v", id, err)
		return nil, err
	}

	for rows.Next() {
		association := new(entity.AssociationEntity)
		if err := rows.StructScan(association); err != nil {
			logrus.Debugf("failed to scan association user from db record: %v", err)
			return nil, err
		}
		return association, nil
	}
	return nil, fmt.Errorf("could not find association user with id=%d", id)
}

func (pdb *PostgresDB) getInvitationById(tx sqlx.Queryer, id int) (*entity.InvitationEntity, error) {
	query := `select * from invites where id = $1`
	rows, err := tx.Queryx(query, id)
	if err != nil {
		logrus.Debugf("failed to get invitation  with id=%d in db: %v", id, err)
		return nil, err
	}

	for rows.Next() {
		inv := new(entity.InvitationEntity)
		if err := rows.StructScan(inv); err != nil {
			logrus.Debugf("failed to scan invitation user from db record: %v", err)
			return nil, err
		}
		return inv, nil
	}
	return nil, fmt.Errorf("could not find invitation user with id=%d", id)

}

func (pdb *PostgresDB) createAssociation(tx NamedQuerier, association *entity.AssociationEntity) (int, error) {
	query := `insert into associations
		(
			name,
			website_url,
			focus
		)
		values (
			:name,
			:website_url,
			:focus
		)
		returning id`
	associationId, err := PreparedQuery(tx, query, association)
	if err != nil {
		logrus.Debugf("failed to insert association in db: %v", err)
		return 0, err
	}
	return associationId, nil
}

func (pdb *PostgresDB) createInvitation(tx NamedQuerier, inv *entity.InvitationEntity) (int, error) {
	query := `insert into invites
		(
		 	creator_id,
			kind,
			role,
		 	entity_id,
			email,
		 	expire_at
		)
		values (
		 	:creator_id,
			:kind,
			:role,
		 	:entity_id,
			:email,
		 	:expire_at
		)
		returning id`
	invId, err := PreparedQuery(tx, query, inv)
	if err != nil {
		logrus.Debugf("failed to insert invitations in db: %v", err)
		return 0, err
	}
	return invId, nil
}

func (pdb *PostgresDB) CreateAssociationUser(associationUser *entity.AssociationUserEntity) (*entity.AssociationUserEntity, error) {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	userId, err := pdb.createUser(tx, associationUser.UserEntity)
	if err != nil {
		return nil, err
	}
	associationUser.UserID = userId

	query := `insert into association_users (user_id, association_id, role)
				values (:user_id, :association_id, :role)
				returning id`
	associationUserId, err := PreparedQuery(tx, query, associationUser)
	if err != nil {
		logrus.Debugf("failed to insert association user in db: %v", err)
		return nil, err
	}
	res, err := pdb.getAssociationUserById(tx, associationUserId)
	if err != nil {
		logrus.Errorf("getting added association user from db: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit association user creation in db: %v", err)
		return nil, err
	}
	return res, nil
}

func (pdb *PostgresDB) EditAssociationUser(id int, associationUser *entity.AssociationUserEntity) (*entity.AssociationUserEntity, error) {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	userId, err := pdb.editUser(tx, id, associationUser.UserEntity)
	if err != nil {
		return nil, err
	}
	associationUser.UserID = userId

	query := `update association_users
				set association_id=:association_id
				where user_id=:user_id
				returning id`
	associationUserId, err := PreparedQuery(tx, query, associationUser)
	if err != nil {
		logrus.Debugf("failed to edit association user in db: %v", err)
		return nil, err
	}
	res, err := pdb.getAssociationUserById(tx, associationUserId)
	if err != nil {
		logrus.Errorf("getting edited association user from db: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit association user update in db: %v", err)
		return nil, err
	}
	return res, nil
}

func (pdb *PostgresDB) CreateCandidate(candidate *entity.CandidateEntity) (*entity.CandidateEntity, error) {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	userId, err := pdb.createUser(tx, candidate.UserEntity)
	if err != nil {
		return nil, err
	}
	candidate.UserID = userId

	query := `insert into candidates
				(
				 	user_id,
					years_of_experience,
					job_status,
					seek_job_type,
					seek_company_size,
					seek_location_type,
					seek_salary,
					seek_values,
					work_permit,
					notice_period
				)
				values (
					:user_id,
					:years_of_experience,
					:job_status,
					:seek_job_type,
					:seek_company_size,
					:seek_location_type,
					:seek_salary,
					:seek_values,
					:work_permit,
					:notice_period
				)
				returning id`
	candidateId, err := PreparedQuery(tx, query, candidate)
	if err != nil {
		logrus.Debugf("failed to insert candidate in db: %v", err)
		return nil, err
	}
	res, err := pdb.getCandidateById(tx, candidateId)
	if err != nil {
		logrus.Errorf("getting added candidate from db: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit candidate creation in db: %v", err)
		return nil, err
	}
	return res, nil
}

func (pdb *PostgresDB) EditCandidate(id int, candidate *entity.CandidateEntity) (*entity.CandidateEntity, error) {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	userId, err := pdb.editUser(tx, id, candidate.UserEntity)
	if err != nil {
		return nil, err
	}
	candidate.UserID = userId

	query := `update candidates
				set user_id=:user_id,
					years_of_experience=:years_of_experience,
					job_status=:job_status,
					seek_job_type=:seek_job_type,
					seek_company_size=:seek_company_size,
					seek_location_type=:seek_location_type,
					seek_salary=:seek_salary,
					seek_values=:seek_values,
					work_permit=:work_permit,
					notice_period=:notice_period
				where user_id=:user_id
				returning id`
	candidateId, err := PreparedQuery(tx, query, candidate)
	if err != nil {
		logrus.Debugf("failed to edit candidate in db: %v", err)
		return nil, err
	}
	res, err := pdb.getCandidateById(tx, candidateId)
	if err != nil {
		logrus.Errorf("getting edited candidate from db: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit candidate update in db: %v", err)
		return nil, err
	}
	return res, nil
}

func (pdb *PostgresDB) CreateCompanyUser(companyUser *entity.CompanyUserEntity) (*entity.CompanyUserEntity, error) {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	userId, err := pdb.createUser(tx, companyUser.UserEntity)
	if err != nil {
		return nil, err
	}
	companyUser.UserID = userId

	query := `insert into company_users (user_id, company_id, role)
				values (:user_id, :company_id, :role)
				returning id`
	companyUserId, err := PreparedQuery(tx, query, companyUser)
	if err != nil {
		logrus.Debugf("failed to insert company user in db: %v", err)
		return nil, err
	}
	res, err := pdb.getCompanyUserById(tx, companyUserId)
	if err != nil {
		logrus.Errorf("getting added company user from db: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit company user creation in db: %v", err)
		return nil, err
	}
	return res, nil
}

func (pdb *PostgresDB) EditCompanyUser(id int, companyUser *entity.CompanyUserEntity) (*entity.CompanyUserEntity, error) {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()

	userId, err := pdb.editUser(tx, id, companyUser.UserEntity)
	if err != nil {
		return nil, err
	}
	companyUser.UserID = userId

	query := `update company_users
				set company_id=:company_id
				where user_id=:user_id
				returning id`
	companyUserId, err := PreparedQuery(tx, query, companyUser)
	if err != nil {
		logrus.Debugf("failed to edit company user in db: %v", err)
		return nil, err
	}
	res, err := pdb.getCompanyUserById(tx, companyUserId)
	if err != nil {
		logrus.Errorf("getting edited company user from db: %v", err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		logrus.Errorf("failed to commit company user creation in db: %v", err)
		return nil, err
	}
	return res, nil
}

func (pdb *PostgresDB) AssignUserPhoto(record *entity.UserPhotoEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteUserPhoto(tx, record.UserID); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertUserPhoto(tx, record); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) DeleteUserPhoto(userId int) error {
	return pdb.deleteUserPhoto(pdb.db, userId)
}

func (pdb *PostgresDB) GetCandidateSkills(candidateId int) (entity.CandidateSkillsEntity, error) {
	res := make(entity.CandidateSkillsEntity, 0)
	query := `select * from candidate_skills where candidate_id = $1`

	rows, err := pdb.db.Queryx(query, candidateId)
	if err != nil {
		return nil, fmt.Errorf("fetching skills for candidate in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.CandidateSkillEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan candidate skills from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) AssignCandidateSkills(candidateId int, records entity.CandidateSkillsEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteCandidateSkills(tx, candidateId); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertCandidateSkills(tx, records); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (pdb *PostgresDB) DeleteCandidateSkills(candidateId int) error {
	return pdb.deleteCandidateSkills(pdb.db, candidateId)
}

func (pdb *PostgresDB) GetCandidateSpokenLanguages(candidateId int) (entity.CandidateSpokenLanguagesEntity, error) {
	res := make(entity.CandidateSpokenLanguagesEntity, 0)
	query := `select * from candidate_spoken_languages where candidate_id = $1`

	rows, err := pdb.db.Queryx(query, candidateId)
	if err != nil {
		return nil, fmt.Errorf("fetching spoken languages for candidate in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.CandidateSpokenLanguageEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan candidate spoken languages from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) AssignCandidateSpokenLanguages(candidateId int, records entity.CandidateSpokenLanguagesEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteCandidateSpokenLanguages(tx, candidateId); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertCandidateSpokenLanguages(tx, records); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (pdb *PostgresDB) DeleteCandidateSpokenLanguages(candidateId int) error {
	return pdb.deleteCandidateSpokenLanguages(pdb.db, candidateId)
}

func (pdb *PostgresDB) GetCandidateSeekLocations(candidateId int) (entity.CandidateSeekLocationsEntity, error) {
	res := make(entity.CandidateSeekLocationsEntity, 0)
	query := `select * from candidate_seek_locations where candidate_id = $1`

	rows, err := pdb.db.Queryx(query, candidateId)
	if err != nil {
		return nil, fmt.Errorf("fetching seek locations for candidate in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.CandidateSeekLocationEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan candidate seek locations from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) AssignCandidateSeekLocations(candidateId int, records entity.CandidateSeekLocationsEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteCandidateSeekLocations(tx, candidateId); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertCandidateSeekLocations(tx, records); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (pdb *PostgresDB) DeleteCandidateSeekLocations(candidateId int) error {
	return pdb.deleteCandidateSeekLocations(pdb.db, candidateId)
}
func (pdb *PostgresDB) AssignCandidateCV(record *entity.CandidateCVEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteCandidateCV(tx, record.CandidateID); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertCandidateCV(tx, record); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (pdb *PostgresDB) DeleteCandidateCV(candidateId int) error {
	return pdb.deleteCandidateCV(pdb.db, candidateId)
}

func (pdb *PostgresDB) GetCandidateAttachments(candidateId int) (entity.CandidateAttachmentsEntity, error) {
	res := make(entity.CandidateAttachmentsEntity, 0)
	query := `select * from candidate_attachments where candidate_id = $1`

	rows, err := pdb.db.Queryx(query, candidateId)
	if err != nil {
		return nil, fmt.Errorf("fetching attachments for candidate in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.CandidateAttachmentEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan candidate attachments from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) AssignCandidateAttachments(candidateId int, records entity.CandidateAttachmentsEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteCandidateAttachments(tx, candidateId); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertCandidateAttachments(tx, records); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (pdb *PostgresDB) DeleteCandidateAttachments(candidateId int) error {
	return pdb.deleteCandidateAttachments(pdb.db, candidateId)
}
func (pdb *PostgresDB) AssignCandidateVideo(record *entity.CandidateVideoEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteCandidateVideo(tx, record.CandidateID); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertCandidateVideo(tx, record); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (pdb *PostgresDB) DeleteCandidateVideo(candidateId int) error {
	return pdb.deleteCandidateVideo(pdb.db, candidateId)
}

func (pdb *PostgresDB) GetCandidateEducationHistoryList(candidateId int) (entity.CandidateEducationHistoryListEntity, error) {
	res := make(entity.CandidateEducationHistoryListEntity, 0)
	query := `select * from candidate_education_history where candidate_id = $1`

	rows, err := pdb.db.Queryx(query, candidateId)
	if err != nil {
		return nil, fmt.Errorf("fetching education history for candidate in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.CandidateEducationHistoryEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan candidate education history from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) AssignCandidateEducationHistoryList(candidateId int, records entity.CandidateEducationHistoryListEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteCandidateEducationHistoryList(tx, candidateId); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertCandidateEducationHistoryList(tx, records); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (pdb *PostgresDB) DeleteCandidateEducationHistoryList(candidateId int) error {
	return pdb.deleteCandidateEducationHistoryList(pdb.db, candidateId)
}

func (pdb *PostgresDB) GetCandidateEmploymentHistoryList(candidateId int) (entity.CandidateEmploymentHistoryListEntity, error) {
	res := make(entity.CandidateEmploymentHistoryListEntity, 0)
	query := `select * from candidate_employment_history where candidate_id = $1`

	rows, err := pdb.db.Queryx(query, candidateId)
	if err != nil {
		return nil, fmt.Errorf("fetching employment history for candidate in db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		view := new(entity.CandidateEmploymentHistoryEntity)
		if err := rows.StructScan(view); err != nil {
			logrus.Errorf("failed to scan candidate employment history from db row: %v", err)
			return nil, err
		}
		res = append(res, view)
	}

	return res, nil
}

func (pdb *PostgresDB) AssignCandidateEmploymentHistoryList(candidateId int, records entity.CandidateEmploymentHistoryListEntity) error {
	tx := pdb.db.MustBegin()
	defer tx.Rollback()
	if err := pdb.deleteCandidateEmploymentHistoryList(tx, candidateId); err != nil {
		return fmt.Errorf("deleting previous data: %v", err)
	}
	if err := pdb.insertCandidateEmploymentHistoryList(tx, records); err != nil {
		return fmt.Errorf("inserting new data: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (pdb *PostgresDB) DeleteCandidateEmploymentHistoryList(candidateId int) error {
	return pdb.deleteCandidateEmploymentHistoryList(pdb.db, candidateId)
}

func (pdb *PostgresDB) insertUserPhoto(tx NamedQuerier, record *entity.UserPhotoEntity) error {
	query := `insert into user_photos (user_id, image_url) values (:user_id, :image_url)`
	if _, err := tx.NamedExec(query, record); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteUserPhoto(tx sqlx.Execer, userId int) error {
	query := `delete from user_photos where user_id = $1`
	if _, err := tx.Exec(query, userId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) insertCandidateSkills(tx NamedQuerier, records entity.CandidateSkillsEntity) error {
	query := `insert into candidate_skills (candidate_id, name, years) values (:candidate_id, :name, :years)`
	for _, record := range records {
		if _, err := tx.NamedExec(query, record); err != nil {
			return err
		}
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateSkills(tx sqlx.Execer, candidateId int) error {
	query := `delete from candidate_skills where candidate_id = $1`
	if _, err := tx.Exec(query, candidateId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateSkillsByUserId(tx sqlx.Execer, userId int) error {
	query := `delete from candidate_skills using candidates where candidates.user_id = $1`
	if _, err := tx.Exec(query, userId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) insertCandidateSpokenLanguages(tx NamedQuerier, records entity.CandidateSpokenLanguagesEntity) error {
	query := `insert into candidate_spoken_languages (candidate_id, language_id, language_name, language_short_name, level) values (:candidate_id, :language_id, :language_name, :language_short_name, :level)`
	for _, record := range records {
		if _, err := tx.NamedExec(query, record); err != nil {
			return err
		}
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateSpokenLanguages(tx sqlx.Execer, candidateId int) error {
	query := `delete from candidate_spoken_languages where candidate_id = $1`
	if _, err := tx.Exec(query, candidateId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateSpokenLanguagesByUserId(tx sqlx.Execer, userId int) error {
	query := `delete from candidate_spoken_languages using candidates where candidates.user_id = $1`
	if _, err := tx.Exec(query, userId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) insertCandidateSeekLocations(tx NamedQuerier, records entity.CandidateSeekLocationsEntity) error {
	query := `insert into candidate_seek_locations (candidate_id, city_id, city_name) values (:candidate_id, :city_id, :city_name)`
	for _, record := range records {
		if _, err := tx.NamedExec(query, record); err != nil {
			return err
		}
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateSeekLocations(tx sqlx.Execer, candidateId int) error {
	query := `delete from candidate_seek_locations where candidate_id = $1`
	if _, err := tx.Exec(query, candidateId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateSeekLocationsByUserId(tx sqlx.Execer, userId int) error {
	query := `delete from candidate_seek_locations using candidates where candidates.user_id = $1`
	if _, err := tx.Exec(query, userId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) insertCandidateCV(tx NamedQuerier, record *entity.CandidateCVEntity) error {
	query := `insert into candidate_cvs (candidate_id, cv_url) values (:candidate_id, :cv_url)`
	if _, err := tx.NamedExec(query, record); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateCV(tx sqlx.Execer, candidateId int) error {
	query := `delete from candidate_cvs where candidate_id = $1`
	if _, err := tx.Exec(query, candidateId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateCVByUserId(tx sqlx.Execer, userId int) error {
	query := `delete from candidate_cvs using candidates where candidates.user_id = $1`
	if _, err := tx.Exec(query, userId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) insertCandidateAttachments(tx NamedQuerier, records entity.CandidateAttachmentsEntity) error {
	query := `insert into candidate_attachments (candidate_id, attachment_url) values (:candidate_id, :attachment_url)`
	for _, record := range records {
		if _, err := tx.NamedExec(query, record); err != nil {
			return err
		}
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateAttachments(tx sqlx.Execer, candidateId int) error {
	query := `delete from candidate_attachments where candidate_id = $1`
	if _, err := tx.Exec(query, candidateId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateAttachmentsByUserId(tx sqlx.Execer, userId int) error {
	query := `delete from candidate_attachments using candidates where candidates.user_id = $1`
	if _, err := tx.Exec(query, userId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) insertCandidateVideo(tx NamedQuerier, record *entity.CandidateVideoEntity) error {
	query := `insert into candidate_videos (candidate_id, video_url) values (:candidate_id, :video_url)`
	if _, err := tx.NamedExec(query, record); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateVideo(tx sqlx.Execer, candidateId int) error {
	query := `delete from candidate_videos where candidate_id = $1`
	if _, err := tx.Exec(query, candidateId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateVideoByUserId(tx sqlx.Execer, userId int) error {
	query := `delete from candidate_videos using candidates where candidates.user_id = $1`
	if _, err := tx.Exec(query, userId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) insertCandidateEducationHistoryList(tx NamedQuerier, records entity.CandidateEducationHistoryListEntity) error {
	query := `insert into candidate_education_history (candidate_id, title, description, entity, from_date, to_date) values (:candidate_id, :title, :description, :entity, :from_date, :to_date)`
	for _, record := range records {
		if _, err := tx.NamedExec(query, record); err != nil {
			return err
		}
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateEducationHistoryList(tx sqlx.Execer, candidateId int) error {
	query := `delete from candidate_education_history where candidate_id = $1`
	if _, err := tx.Exec(query, candidateId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateEducationHistoryListByUserId(tx sqlx.Execer, userId int) error {
	query := `delete from candidate_education_history using candidates where candidates.user_id = $1`
	if _, err := tx.Exec(query, userId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) insertCandidateEmploymentHistoryList(tx NamedQuerier, records entity.CandidateEmploymentHistoryListEntity) error {
	query := `insert into candidate_employment_history (candidate_id, title, description, company, from_date, to_date) values (:candidate_id, :title, :description, :company, :from_date, :to_date)`
	for _, record := range records {
		if _, err := tx.NamedExec(query, record); err != nil {
			return err
		}
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateEmploymentHistoryList(tx sqlx.Execer, candidateId int) error {
	query := `delete from candidate_employment_history where candidate_id = $1`
	if _, err := tx.Exec(query, candidateId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) deleteCandidateEmploymentHistoryListByUserId(tx sqlx.Execer, userId int) error {
	query := `delete from candidate_employment_history using candidates where candidates.user_id = $1`
	if _, err := tx.Exec(query, userId); err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDB) getUserById(tx sqlx.Queryer, id int) (*entity.UserEntity, error) {
	query := `select * from users where id = $1`
	rows, err := tx.Queryx(query, id)
	if err != nil {
		logrus.Debugf("failed to get user with id=%d in db: %v", id, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := new(entity.UserEntity)
		if err := rows.StructScan(user); err != nil {
			logrus.Errorf("failed to scan user from db row: %v", err)
			return nil, err
		}
		return user, nil
	}
	return nil, fmt.Errorf("could not find user with id=%d", id)
}

func (pdb *PostgresDB) getAssociationUserById(tx sqlx.Queryer, id int) (*entity.AssociationUserEntity, error) {
	query := `select users.*, association_users.*
				from users
				inner join association_users on users.id = association_users.user_id
				where association_users.id = $1`
	rows, err := tx.Queryx(query, id)
	if err != nil {
		logrus.Debugf("failed to get association user with id=%d in db: %v", id, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		associationUser := new(entity.AssociationUserEntity)
		if err := rows.StructScan(associationUser); err != nil {
			logrus.Errorf("failed to scan association user from db row: %v", err)
			return nil, err
		}
		return associationUser, nil
	}
	return nil, fmt.Errorf("could not find association user with id=%d", id)
}

func (pdb *PostgresDB) getCandidateById(tx sqlx.Queryer, id int) (*entity.CandidateEntity, error) {
	query := `select users.*, candidates.*
				from users
				inner join candidates on users.id = candidates.user_id
				where candidates.id = $1`
	rows, err := tx.Queryx(query, id)
	if err != nil {
		logrus.Debugf("failed to get candidate with id=%d in db: %v", id, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		candidate := new(entity.CandidateEntity)
		if err := rows.StructScan(candidate); err != nil {
			logrus.Errorf("failed to scan candidate from db row: %v", err)
			return nil, err
		}
		return candidate, nil
	}
	return nil, fmt.Errorf("could not find candidate with id=%d", id)
}

func (pdb *PostgresDB) getCompanyUserById(tx sqlx.Queryer, id int) (*entity.CompanyUserEntity, error) {
	query := `select users.*, company_users.*
				from users
				inner join company_users on users.id = company_users.user_id
				where company_users.id = $1`
	rows, err := tx.Queryx(query, id)
	if err != nil {
		logrus.Debugf("failed to get company user with id=%d in db: %v", id, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		companyUser := new(entity.CompanyUserEntity)
		if err := rows.StructScan(companyUser); err != nil {
			logrus.Errorf("failed to scan company user from db row: %v", err)
			return nil, err
		}
		return companyUser, nil
	}
	return nil, fmt.Errorf("could not find company user with id=%d", id)
}

func (pdb *PostgresDB) createUser(tx NamedQuerier, user *entity.UserEntity) (int, error) {
	query := `insert into users
				(
				 	kind,
					first_name,
					last_name,
					preferred_name,
					email,
					phone_number,
					birth_date,
					linkedin_url,
					github_url,
					portfolio_url
				)
				values (
					:kind,
					:first_name,
					:last_name,
					:preferred_name,
					:email,
					:phone_number,
					:birth_date,
					:linkedin_url,
					:github_url,
					:portfolio_url
				)
				returning id`
	userId, err := PreparedQuery(tx, query, user)
	if err != nil {
		logrus.Debugf("failed to insert user in db: %v", err)
		return 0, err
	}
	return userId, nil
}

func (pdb *PostgresDB) editUser(tx NamedQuerier, id int, user *entity.UserEntity) (int, error) {
	user.ID = id
	query := `update users
				set first_name=:first_name,
					last_name=:last_name,
					preferred_name=:preferred_name,
					email=:email,
					phone_number=:phone_number,
					birth_date=:birth_date,
					linkedin_url=:linkedin_url,
					github_url=:github_url,
					portfolio_url=:portfolio_url
				where id=:id
				returning id`
	userId, err := PreparedQuery(tx, query, user)
	if err != nil {
		logrus.Debugf("failed to update user in db: %v", err)
		return 0, err
	}
	return userId, nil
}
