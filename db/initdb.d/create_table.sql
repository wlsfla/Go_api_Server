CREATE TABLE GoAPIService.update_info (
    host_ip VARCHAR(45) NOT NULL,
    host_name VARCHAR(255),
    winver VARCHAR(10),
    buildver VARCHAR(45),
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    result TINYINT,
    PRIMARY KEY (host_ip)
);

CREATE TABLE GoAPIService.target_winver (
	winver VARCHAR(10),
    buildver VARCHAR(45)
);

insert into	GoAPIService.target_winver
values
	("1803", "17134.2208"),
	("1809", "17763.3653"),
	("1903", "18362.1256"),
	("1909", "18363.2274"),
	("20H2", "19042.2311"),
	("21H1", "19043.2311"),
	("21H2", "19044.2311");

