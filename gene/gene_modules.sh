#!/bin/sh
#go run . -ro ../modules $1
rootb="D:/Dev/tmp/go/pkg/mod/github.com/larksuite/base-sdk-go/v3@v3.0.2/"
root="D:/Dev/tmp/go/pkg/mod/github.com/larksuite/oapi-sdk-go/v3@v3.3.0/"

./gene -ro ../modules $rootb/
./gene -ro ../modules $rootb/core
./gene -ro ../modules $rootb/service/base/v1 # base sdk
./gene -ro ../modules $rootb/service/drive/v1 # base sdk

./gene -ro ../modules $root/
./gene -i=github.com/gogo/protobuf/proto -ro ../modules $root/ws
./gene -ro ../modules $root/event
./gene -ro ../modules $root/event/dispatcher
./gene -ro ../modules $root/cache
./gene -ro ../modules $root/card
./gene -ro ../modules $root/core/httpserverext

./gene -ro ../modules $root/service/wiki
./gene -ro ../modules $root/service/wiki/v2
./gene -ro ../modules $root/service/workplace
./gene -ro ../modules $root/service/workplace/v1
./gene -ro ../modules $root/service/verification
./gene -ro ../modules $root/service/verification/v1
./gene -ro ../modules $root/service/vc
./gene -ro ../modules $root/service/vc/v1
./gene -ro ../modules $root/service/translation
./gene -ro ../modules $root/service/translation/v1
./gene -ro ../modules $root/service/compensation
./gene -ro ../modules $root/service/compensation/v1
./gene -ro ../modules $root/service/tenant
./gene -ro ../modules $root/service/tenant/v2
./gene -ro ../modules $root/service/task
./gene -ro ../modules $root/service/task/v1
./gene -ro ../modules $root/service/task/v2
./gene -ro ../modules $root/service/speech_to_text
./gene -ro ../modules $root/service/speech_to_text/v1
./gene -ro ../modules $root/service/sheets
./gene -ro ../modules $root/service/sheets/v3
./gene -ro ../modules $root/service/security_and_compliance
./gene -ro ../modules $root/service/security_and_compliance/v1
./gene -ro ../modules $root/service/report
./gene -ro ../modules $root/service/report/v1
./gene -ro ../modules $root/service/search
./gene -ro ../modules $root/service/search/v2
./gene -ro ../modules $root/service/personal_settings
./gene -ro ../modules $root/service/personal_settings/v1
./gene -ro ../modules $root/service/passport
./gene -ro ../modules $root/service/passport/v1
./gene -ro ../modules $root/service/optical_char_recognition
./gene -ro ../modules $root/service/optical_char_recognition/v1
./gene -ro ../modules $root/service/okr
./gene -ro ../modules $root/service/okr/v1
./gene -ro ../modules $root/service/meeting_room
./gene -ro ../modules $root/service/meeting_room/v1
./gene -ro ../modules $root/service/mdm
./gene -ro ../modules $root/service/mdm/v1
./gene -ro ../modules $root/service/mail
./gene -ro ../modules $root/service/mail/v1
./gene -ro ../modules $root/service/lingo
./gene -ro ../modules $root/service/lingo/v1
./gene -ro ../modules $root/service/im
./gene -ro ../modules $root/service/im/v1
./gene -ro ../modules $root/service/im/v2
./gene -ro ../modules $root/service/human_authentication
./gene -ro ../modules $root/service/human_authentication/v1
./gene -ro ../modules $root/service/hire
./gene -ro ../modules $root/service/hire/v1
./gene -ro ../modules $root/service/helpdesk
./gene -ro ../modules $root/service/helpdesk/v1
./gene -ro ../modules $root/service/ext
#./gene -ro ../modules $root/service/face_detection empty
./gene -ro ../modules $root/service/face_detection/v1
./gene -ro ../modules $root/service/gray_test_open_sg
./gene -ro ../modules $root/service/gray_test_open_sg/v1
./gene -ro ../modules $root/service/event
./gene -ro ../modules $root/service/event/v1
./gene -ro ../modules $root/service/drive
./gene -ro ../modules $root/service/drive/v1
./gene -ro ../modules $root/service/drive/v2
./gene -ro ../modules $root/service/ehr
./gene -ro ../modules $root/service/ehr/v1
./gene -ro ../modules $root/service/docx
./gene -ro ../modules $root/service/docx/v1
./gene -ro ../modules $root/service/document_ai
./gene -ro ../modules $root/service/document_ai/v1
./gene -ro ../modules $root/service/corehr
./gene -ro ../modules $root/service/corehr/v1
./gene -ro ../modules $root/service/corehr/v2
./gene -ro ../modules $root/service/contact
./gene -ro ../modules $root/service/contact/v3
./gene -ro ../modules $root/service/calendar
./gene -ro ../modules $root/service/calendar/v4
./gene -ro ../modules $root/service/block
./gene -ro ../modules $root/service/block/v2
./gene -ro ../modules $root/service/board
./gene -ro ../modules $root/service/board/v1
./gene -ro ../modules $root/service/bitable
./gene -ro ../modules $root/service/bitable/v1
./gene -ro ../modules $root/service/baike
./gene -ro ../modules $root/service/baike/v1
./gene -ro ../modules $root/service/authen
./gene -ro ../modules $root/service/authen/v1
./gene -ro ../modules $root/service/auth
./gene -ro ../modules $root/service/auth/v3
./gene -ro ../modules $root/service/attendance
./gene -ro ../modules $root/service/attendance/v1
./gene -ro ../modules $root/service/approval
./gene -ro ../modules $root/service/approval/v4
./gene -ro ../modules $root/service/application
./gene -ro ../modules $root/service/application/v6
./gene -ro ../modules $root/service/aily
./gene -ro ../modules $root/service/aily/v1
./gene -ro ../modules $root/service/admin
./gene -ro ../modules $root/service/admin/v1
./gene -ro ../modules $root/service/acs
./gene -ro ../modules $root/service/acs/v1

