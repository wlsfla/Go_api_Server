
## A. 개요
---
	windows desktop 대상으로 각각 version 별 업데이트 파일을 다르게 하고 결과 log를 종합하는 porject

## B. 구성 요소
---
	가. Client: Cmd 실행 - ps1 파일 다운로드/실행
	나. Api Server: 실제 데이터 처리의 중심이 되는 Gateway.
		- C-가. 세부 업데이트 프로세스 로직이 처리되는 파일을 제공한다. 파일에 문제가 생겼을 경우 서버측 파일만 바꿔주고 Client는 cmd 파일만 다시 실행하면 된다.
		- C-나, C-다. input value에 대한 DB insert를 수행.
	다. File Server: 제공하고자 하는 파일 url이 저장된 서버
	라. DB: Api Server로부터 들어온 값을 저장.(C-나, C-다. 참고)
	
## C. Api Description.
---
	가. ps1 file request. @
		* iuput: [Get] /update/ps
		* output:
	    ```
			return powershell file
        ```
	
	나. info register @
		* iuput: [post] /api/info_reg/:hostname/:winver
		* output:
		```
        	{
				data:{
					url:{
						{:url1},{:url2},{:url3} ...
					}
				}
				state:1
			}
        ```
		
	
	다. Result Report @
		* iuput: [post] /api/result/:hostname/:result
		* output: None

## D. DB Description
---
	가. Attribute
        - pc_ip: Sender Client ip 사용
        - hostname: hostname
        - created_time: yyyy-mm-dd hh:mm:ss
        - updated_time: yyyy-mm-dd hh:mm:ss
        - result: 0: 미수행, 1: 성공, 2: 오류 3: 해당없음

## E. Update Version Check Logic
---
if (현재 시스템의 winver이 1803이면)
    return 수동 업데이트
if (현재 시스템의 빌드버전이 같으면)
    - 적용 완료
if (현재 시스템의 빌드버전이 높으면)
    - 업데이트 불가(기 업데이트 완료)

#현재 시스템의 빌드버전이 낮으면
if (업데이트 목록에 현재 업데이트가 있으면)
        - 적용 완료

업데이트 스케줄러 등록