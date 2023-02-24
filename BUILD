load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "nogo")

nogo(
    name = "vet",
    vet = True,
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_x_tools//go/analysis/passes/asmdecl:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/assign:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/atomicalign:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/cgocall:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/composite:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/copylock:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/deepequalerrors:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/errorsas:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/fieldalignment:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/framepointer:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/httpresponse:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/ifaceassert:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/loopclosure:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/lostcancel:go_default_library",
        # "@org_golang_x_tools//go/analysis/passes/nilness:go_default_library", # template methods currently cause this analyzer to panic
        "@org_golang_x_tools//go/analysis/passes/shadow:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/shift:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/sortslice:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/stdmethods:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/stringintconv:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/structtag:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/tests:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/testinggoroutine:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/unmarshal:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/unreachable:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/unsafeptr:go_default_library",
        "@org_golang_x_tools//go/analysis/passes/unusedresult:go_default_library",
    ],
)

# Prefer generated BUILD files to be called BUILD over BUILD.bazel
# gazelle:build_file_name BUILD,BUILD.bazel
# gazelle:prefix github.com/luluz66/godemo
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_library(
    name = "godemo_lib",
    srcs = ["main.go"],
    importpath = "github.com/luluz66/godemo",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "godemo",
    embed = [":godemo_lib"],
    visibility = ["//visibility:public"],
)
