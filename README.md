
## A. 개요
---
windows desktop 대상 각각 version 별 업데이트 파일을 다르게 하고 결과 log를 종합하는 porject

## B. 구성 요소
---
- Client: Cmd 실행 - ps1 파일 다운로드/실행
- Api Server: 실제 데이터 처리의 중심이 되는 Gateway.
- C-가. 세부 업데이트 프로세스 로직이 처리되는 파일을 제공한다. 파일에 문제가 생겼을 경우 서버측 파일만 바꿔주고 Client는 cmd 파일만 다시 실행하면 된다.
- C-나, C-다. input value에 대한 DB insert를 수행.
. File Server: 제공하고자 하는 파일 url이 저장된 서버
. DB: Api Server로부터 들어온 값을 저장.(C-나, C-다. 참고)
	
## C. Api Description.
---
- ps1 file request.
	* iuput: `[get]` /update/ps
	* output:
	    ```
			return powershell file
        ```
	
- info register
	* iuput: `[post]` /api/info_reg/`:hostname`/`:winver`
	* output:
		```json
		{
			data:{
				url:{
					{:url1},{:url2},{:url3} ...
				}
			}
			state:1
		}
        ```
	
- Result Report
	* iuput: `[post]` /api/result/`:result`
	* output: `None`

## D. DB Description
---
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
---

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
---
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