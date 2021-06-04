WHOAMI=$(shell whoami)

SED_I = sed -i
ifeq ($(shell uname -s), Darwin)
    SED_I = sed -i ''
endif

nothing:

generate_proto:
	protoc -I. --go_out=. --go_opt=Mgoogle_libphonenumber/resources/phonemetadata.proto=../libphonenumber --go_opt=Mgoogle_libphonenumber/resources/phonenumber.proto=../libphonenumber ./google_libphonenumber/resources/*.proto
	awk '/static const unsigned char/ { show=1 } show; /}/ { show=0 }' ./google_libphonenumber/cpp/src/phonenumbers/metadata.cc | tail -n +2 | sed '$$d' | sed -E 's/([^,])$$/\1,/g' | awk 'BEGIN{print "package libphonenumber\nvar metaData = []byte{"}; {print}; END{print "}"}' > metagen.go
	go fmt ./metagen.go

distupdate:
	rm -rf ./google_libphonenumber
	git clone --depth 1 https://github.com/googlei18n/libphonenumber.git ./google_libphonenumber/

update: distupdate generate_proto
