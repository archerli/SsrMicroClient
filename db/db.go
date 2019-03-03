package main
import(
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main(){
	db,err := sql.Open("sqlite3","./test.db")
	if err!=nil{
		fmt.Println(err)
		return
	}


	//清空表
	db.Exec("DELETE FROM SSR_info")


    //删除表
	db.Exec("DROP TABLE IF EXISTS SSR_info;")
	
	 //创建表 
	sql_table := `
	CREATE TABLE IF NOT EXISTS SSR_info(
		server TEXT,
		server_port TEXT,
		protocol TEXT,
		method TEXT,
		obfs TEXT,
		password TEXT,
		obfsparam TEXT,
		protoparam TEXT
		);
		`
	db.Exec(sql_table)

	//插入
	stmt,_ := db.Prepare("INSERT INTO SSR_info(server,server_port,protocol,method,obfs,password,obfsparam,protoparam)values(?,?,?,?,?,?,?,?)")
	res,_ := stmt.Exec("x","x","x","x","x","x","x","x")
	id,_ := res.LastInsertId()

	//查找
	rows, err := db.Query("SELECT * FROM SSR_info WHERE server = 'x'")
	var server,server_port,protocol,method,obfs,password,obfsparam,protoparam string
	for rows.Next(){
		err = rows.Scan(&server,&server_port,&protocol,&method,&obfs,&password,&obfsparam,&protoparam)
		fmt.Println(server_port)
	}

	fmt.Println(id)
}