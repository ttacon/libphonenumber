load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

# gazelle:prefix github.com/burstsms/libphonenumber
# gazelle:build_file_name BUILD
gazelle(name = "gazelle")

go_library(
    name = "libphonenumber",
    srcs = [
        "countryCodeToTimeZones.go",
        "countrycodetoregionmap.go",
        "metagen.go",
        "phonemetadata.pb.go",
        "phonenumber.pb.go",
        "phonenumbermatcher.go",
        "phonenumberutil.go",
    ],
    importpath = "github.com/burstsms/libphonenumber",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_protobuf//proto:go_default_library",
        "@github_com_ttacon_builder//:builder",
    ],
)

go_test(
    name = "libphonenumber_test",
    srcs = ["phonenumberutil_test.go"],
    embed = [":libphonenumber"],
    deps = ["@com_github_golang_protobuf//proto:go_default_library"],
)
