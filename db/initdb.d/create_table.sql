CREATE TABLE GoAPIService.update_logs (
    host_ip VARCHAR(45) NOT NULL,
    host_name VARCHAR(255),
    winver VARCHAR(10),
    buildver VARCHAR(45),
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    result TINYINT
);
/*
    select host_ip, host_name, winver, buildver, max(updated_time), result from GoAPIService.update_log group by update_log.host_ip;
    insert into GoAPIService.update_log values ("2.1.1.2", "test_name", "winver", "buildver", default, 0);
*/

CREATE TABLE GoAPIService.target_winvers (
	winver VARCHAR(10),
    buildver VARCHAR(45),
    kbNumber VARCHAR(45)
);

insert into	GoAPIService.target_winvers
values
	("1803", "17134.2208", "KB5003174"),
	("1809", "17763.3653", "KB5021655"),
	("1903", "18362.1256", "KB4592449"),
	("1909", "18363.2274", "KB5013945"),
	("2004", "19041.1415", "KB2008212"),
	("20H2", "19042.2311", "KB5020030"),
	("21H1", "19043.2311", "KB5020030"),
	("21H2", "19044.2311", "KB5020030");

