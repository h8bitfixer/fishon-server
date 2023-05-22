package db

//
//import (
//	"fmt"
//
//	"gorm.io/gorm/schema"
//
//	// "github.com/go-sql-driver/mysql"
//	"sync"
//	"time"
//
//	// "github.com/jinzhu/gorm"
//	// _ "github.com/jinzhu/gorm/dialects/mysql"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//type mysqlDB struct {
//	sync.RWMutex
//	dbMap map[string]*gorm.DB
//}
//
//func initMysqlDB() {
//	// When there is no open IM database, connect to the mysql built-in database to create openIM database
//	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
//		config.Config.Mysql.DBUserName, config.Config.Mysql.DBPassword, config.Config.Mysql.DBAddress[0], "mysql")
//	var db *gorm.DB
//	var err1 error
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: log.GetSqlLogger(constant.MySQLLogFileName), NamingStrategy: schema.NamingStrategy{SingularTable: true}})
//	if err != nil {
//		fmt.Println("0", "Open failed ", err.Error(), dsn)
//		time.Sleep(time.Duration(30) * time.Second)
//		db, err1 = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: log.GetSqlLogger(constant.MySQLLogFileName), NamingStrategy: schema.NamingStrategy{SingularTable: true}})
//		if err1 != nil {
//			fmt.Println("0", "Open failed ", err1.Error(), dsn)
//			panic(err1.Error())
//		}
//	}
//
//	// Check the database and table during initialization
//	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default charset utf8mb4 COLLATE utf8mb4_unicode_ci;", config.Config.Mysql.DBDatabaseName)
//	err = db.Exec(sql).Error
//	if err != nil {
//		fmt.Println("0", "Exec failed ", err.Error(), sql)
//		panic(err.Error())
//	}
//
//	sqlDB, _ := db.DB()
//	sqlDB.Close()
//
//	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
//		config.Config.Mysql.DBUserName, config.Config.Mysql.DBPassword, config.Config.Mysql.DBAddress[0], config.Config.Mysql.DBDatabaseName)
//	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: log.GetSqlLogger(constant.MySQLLogFileName), NamingStrategy: schema.NamingStrategy{SingularTable: true}})
//	if err != nil {
//		fmt.Println("0", "Open failed ", err.Error(), dsn)
//		panic(err.Error())
//	}
//
//	sqlDB, _ = db.DB()
//	fmt.Println("open db ok ", dsn)
//	db.AutoMigrate(
//		&Register{},
//		//&Friend{},
//		//&FriendRequest{},
//		//&Group{},
//		//&GroupMember{},
//		//&GroupRequest{},
//		//&User{}, &PrivacySetting{},
//		//
//		//&Black{}, &BlackForMoment{}, &ChatLog{}, &Register{}, &Conversation{}, &AppVersion{}, &Department{}, &NewAppVersion{}, &DiscoverUrl{},
//		//&InviteCodeLog{}, &InviteCode{}, &InviteCodeRelation{}, &InviteChannelCode{}, &Config{}, &AdminUser{},
//		//&AdminAPIs{}, &AdminPages{}, &AdminRole{}, &MomentSQL{}, &MomentLikeSQL{}, &MomentCommentSQL{}, &OauthClient{},
//		//&Contact{}, &ContactExclude{},
//		//&InterestType{},
//		//&InterestLanguage{},
//		//&InterestUser{},
//		//&InterestGroup{}, &InterestGroupExclude{},
//		//&GroupHeat{},
//		//&Official{}, &OfficialAnalytics{}, &OfficialInterest{},
//		//&OfficialFollowSQL{}, &ArticleLikeSQL{}, &ArticleSQL{}, &ArticleReadSQL{}, &ArticleCommentSQL{}, ArticleCommentLikeSQL{},
//		//&FavoritesSQL{}, &VideoAudioCommunicationRecord{}, &CommunicationGroupMember{},
//		//&GameCategories{}, &GameLink{}, &Game{}, &GamePlayHistory{}, &GameFavorites{},
//		//&MePageURL{},
//		//
//		//&ShortVideo{}, &ShortVideoLike{}, &ShortVideoComment{}, &ShortVideoCommentLike{}, &ShortVideoFollow{}, &ShortVideoUserCount{}, &ShortVideoNotice{},
//	)
//
//	db.Set("gorm:table_options", "CHARSET=utf8mb4")
//	db.Set("gorm:table_options", "collation=utf8mb4_unicode_ci")
//
//	if !db.Migrator().HasTable(&Friend{}) {
//		fmt.Println("CreateTable Friend")
//		db.Migrator().CreateTable(&Friend{})
//	}
//
//	if !db.Migrator().HasTable(&FriendRequest{}) {
//		fmt.Println("CreateTable FriendRequest")
//		db.Migrator().CreateTable(&FriendRequest{})
//	}
//
//	sqlDB.Close()
//}
//
//func (m *mysqlDB) DefaultGormDB() (*gorm.DB, error) {
//	return m.GormDB(config.Config.Mysql.DBAddress[0], config.Config.Mysql.DBDatabaseName)
//}
//
//func (m *mysqlDB) GormDB(dbAddress, dbName string) (*gorm.DB, error) {
//	m.Lock()
//	defer m.Unlock()
//
//	k := key(dbAddress, dbName)
//	if _, ok := m.dbMap[k]; !ok {
//		if err := m.open(dbAddress, dbName); err != nil {
//			return nil, err
//		}
//	}
//	return m.dbMap[k], nil
//}
//
//func (m *mysqlDB) open(dbAddress, dbName string) error {
//	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
//		config.Config.Mysql.DBUserName, config.Config.Mysql.DBPassword, dbAddress, dbName)
//	// db, err := gorm.Open("mysql", dsn, &gorm.Config{Logger: log.GetNewLogger(constant.SQLiteLogFileName)})
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: log.GetSqlLogger(constant.MySQLLogFileName), NamingStrategy: schema.NamingStrategy{SingularTable: true}})
//	if err != nil {
//		return err
//	}
//
//	sqlDB, _ := db.DB()
//	sqlDB.SetMaxOpenConns(config.Config.Mysql.DBMaxOpenConns)
//	sqlDB.SetMaxIdleConns(config.Config.Mysql.DBMaxIdleConns)
//	sqlDB.SetConnMaxLifetime(time.Duration(config.Config.Mysql.DBMaxLifeTime) * time.Second)
//
//	if m.dbMap == nil {
//		m.dbMap = make(map[string]*gorm.DB)
//	}
//	k := key(dbAddress, dbName)
//	m.dbMap[k] = db
//	return nil
//}
