package errcode

var (
	ParseErrorNoBranchMatch      = NewError(1000005, "flow error, no branch match")
	ParseErrorRulesetOutputEmpty = NewError(1000011, "ruleset output is empty")
	ParseErrorTreeNotMatch       = NewError(1000021, "tree not match error")
	ParseErrorTreeOutputEmpty    = NewError(1000022, "tree output is empty")
	ParseErrorMatrixNotMatch     = NewError(1000031, "matrix not match error")
	ParseErrorMatrixOutputEmpty  = NewError(1000032, "matrix output is empty")
	ParseErrorBlockNotMatch      = NewError(1000041, "block not match error")

	ParseErrorNotSupportOperator  = NewError(1000101, "not support operator")
	ParseErrorTargetMustBeArray   = NewError(1000102, "target must be array, check yaml first")
	ParseErrorTargetNotSupport    = NewError(1000103, "target not support error")
	ParseErrorFeatureTypeNotMatch = NewError(1000104, "feature type is not match")
	ParseErrorFeatureSetValue     = NewError(1000105, "the type of date is not match")
)
