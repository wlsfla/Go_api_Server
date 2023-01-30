
## A. 개요
windows desktop 대상 각각 version 별 업데이트 파일을 다르게 하고 결과 log를 종합하는 porject

## B. 구성

	
## C. Api Description.
- Client info Regist
	* input `post` /api/v2/insert/updatelog
		```json
		{
			"Host_ip": "1.1.1.1",
			"Host_name": "dummy_name",
			"Winver": "22H1",
			"Buildver": "19044.2486",
			"Result": 0
		}
		
        ```
	* output
	    ```json
		{
			"host_ip": "1.1.1.1",
			"host_name": "dummy_name",
			"winver": "22H1",
			"buildver": "19044.2486",
			"updated_time": "0001-01-01T00:00:00Z",
			"Result": 0
		}
        ```
	


## D. DB Description
### [table] host_info
|Attr Name|Desc|-|
|-|-|-|
pc_ip|Sender Client ip 사용|-|
hostname|hostname|-|
created_time|yyyy-MM-dd hh:mm:ss|-|
updated_time|yyyy-MM-dd hh:mm:ss|-|
result| `0: 미수행`,`1: 성공`, `2: 오류`, `3: 해당없음`

- `Build Ver` 추가 수집 필요


## E. Update Version Check Logic
- if (현재 시스템의 winver이 1803이면)
	- return 수동 업데이트
- if (현재 시스템의 빌드버전이 같으면)
    - 적용 완료
- if (현재 시스템의 빌드버전이 높으면)
    - 업데이트 불가(기 업데이트 완료)

- if (업데이트 목록에 현재 업데이트가 있으면) //현재 시스템의 빌드버전이 낮으면
        - 적용 완료

- Finaly : 업데이트 스케줄러 등록

## F. update file list
|OS Ver.|KB Update Ver.|Build Ver.(After Update)|
|-|-|-|
|1803.msu|KB5003174|17134.2208|
|1809.msu|KB4592449|18362.1256|
|1909.msu|KB5021655|17763.3653|
|2004.msu|KB5013945|18363.2274|
|20H2.msu|KB5020030|19042.2311|
|21H1.msu|KB5020030|19043.2311|
|21H2.msu|KB5020030|19044.2311|

- 빌드버전의 . (dot)을 기준으로 두 부분 중
	- `앞-뒤` 순서대로 높은 번호일 수록 버전이 높다.
		- (예시) 현재 PC 버전은 21H2(19044.2364)
			- Target 버전보다 빌드 버전이 높고, KB5020030 업데이트가 없다.
			- 실제 업데이트 수행 시 적용 불가하다.




========================================================================


A. Checker
	** 최초 1회 직접 수행.
	** 작업 스케줄러에 등록
		- At windows log on
	
	1) PC 정보 수집
	2) 업데이트 여부 검사
		a) Target Buildver Get
		b) status가 0이면 업데이트 불가
		c) currBuildver is
			Lower Then => result = 1
			Else => result = 9
	3) 점검 결과 전송
	
B. Updater
	업데이트 점검 여부 검사 로직 후
	result가 1이면 업데이트 수행
	
	업데이트 로직
	result == 1 Then
		Downloader
		