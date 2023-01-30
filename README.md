
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



** 윈도우 Client 구성

	A) Task Manager: 이용자가 직접 수행
		1) B),C)를 서버로부터 다운로드. task 등록한다.
			- B) 시간마다 실행
			- C) PC 로그인할 때 실행
		2) C 파일을 실행한다.(파일 유무 검사 필수)
		3) "업데이트 작업 완료." 메시지 박스 시현.
		종료

	B) File Downloader
		1) 서버에 파일 다운로드 가능 API 요청(Get)
		2) status == 1 then url 필드의 addr에 파일 다운로드 시작
		3) status == 0 then.
		종료

	C) pc info checker
		1) pc 정보를 수집한다.
		2) 서버에 insert 요청을 한다.(Post)
		종료
		
=====================================================================

1) 다음 경로로 업데이트 파일 다운로드
	%UserProfile%\polupdate\202212\winupdate.exe
2) 작업 스케줄러 등록
3) 레지트스리 등록
