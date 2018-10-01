package log

const SelectLogShowFilter = "select query,appname,env,ip,create_time,create_user,id,hostname,click from log_show_filter"
const UpdateLogShowFilter = "update log_show_filter"
const InsertLogShowFilter = "insert into log_show_filter" 
const DeleteLogShowFilter = "delete from log_show_filter" 

const SelectLogShowIp = "select id,ip,create_time,app_name from log_show_ip"
const UpdateLogShowIp = "update log_show_ip"
const InsertLogShowIp = "insert into log_show_ip" 
const DeleteLogShowIp = "delete from log_show_ip" 

const SelectLogShowHostname = "select create_time,id,hostname from log_show_hostname"
const UpdateLogShowHostname = "update log_show_hostname"
const InsertLogShowHostname = "insert into log_show_hostname" 
const DeleteLogShowHostname = "delete from log_show_hostname" 

const SelectLogShowAppname = "select id,appname,create_time from log_show_appname"
const UpdateLogShowAppname = "update log_show_appname"
const InsertLogShowAppname = "insert into log_show_appname" 
const DeleteLogShowAppname = "delete from log_show_appname" 

const SelectLogShowHistory = "select env,create_time,create_user,query,appname,hostname,ip,id from log_show_history"
const SelectLastSearch = `select env,create_time,create_user,query,appname,hostname,ip,id from log_show_history where create_user="{0}" order by id desc limit 1`
const UpdateLogShowHistory = "update log_show_history"
const InsertLogShowHistory = "insert into log_show_history" 
const DeleteLogShowHistory = "delete from log_show_history" 
