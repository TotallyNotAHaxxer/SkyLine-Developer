package SkyLine

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////// ***									Error system file for SkyLine 									  *** /////////////////
/////////// 																											 /////////////////
/////////// 																											 /////////////////
/////////// 																											 /////////////////
/////////// 																											 /////////////////
/////////// 																											 /////////////////
/////////// 																											 /////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// EC stands for Error Code

var (
	CODE_PARSE_FLOAT_ERROR                                                       = "EC_100" // Error | Could not parse float value
	CODE_PARSE_INT_ERROR                                                         = "EC_101" // Error | Could not parse int
	CODE_PARSE_STRING_ERROR                                                      = "EC_102" // Error | Could not parse string
	CODE_PARSE_BOOL_ERROR                                                        = "EC_103" // Error | Could not parse bool
	CODE_PARSE_NULL_ERROR                                                        = "EC_104" // Error | Could not parse null
	CODE_PARSE_ARRAY_ERROR                                                       = "EC_105" // Error | Could not parse array
	CODE_PARSE_HASH_ERROR                                                        = "EC_106" // Error | Could not parse hash
	CODE_PARSE_HASHKEY_ERROR                                                     = "EC_107" // Error | Unuseable hash key
	CODE_PARSE_TYPE_ERROR                                                        = "EC_108" // Error | Type mismatch in addition use sprintf to combine all types into a string for
	CODE_PARSE_OPERATOR_ERROR                                                    = "EC_109" // Error | Unknown operator
	CODE_PARSE_IDENTIFIER_ERROR                                                  = "EC_110" // Error | Unknown IDENT / Identifier
	CODE_PARSE_FUNCTION_ARGUMENTS_NOT_ENOUGH_ERROR                               = "EC_111" // Error | Invalid function arguments | Function does not have enough arguments in call to execute the given function
	CODE_PARSE_MACRO_INVALID_ERROR                                               = "EC_112" // Error | Invalid MACRO
	CODE_PARSE_INDEX_OPERATOR_UNSUPPORTED_WITHIN_KEY_NOTE_ERROR                  = "EC_113" // Error | Unsupported index operator must be STRING or INTEGER
	CODE_PARSE_AST_MODIFICATION_TO_MACRO_UNSUPPORTED_METHOD_ERROR                = "EC_114" // Error | Unsupported AST modification method
	CODE_NO_FUNCTIONS_OR_SYMBOLS_LOADED                                          = "EC_115" // Error | No symbols, functions, standards or keywords called, resulted in interpreter violation
	CODE_WRONG_NUMBER_OF_ARGUMENTS                                               = "EC_116" // Error | During call to built in function, length or number of arguments required for the function were not provided
	CODE_PREFIX_PARSE_FUNCTION_INVALID_OR_UNFOUND_WITHIN_PARSER_AND_INTERPRETRR  = "EC_117" // Error | No prefix parse function found for token
	CODE_EXPECT_PEEK_ERROR_DURING_CALL_TO_PEEK                                   = "EC_118" // Error | Expected next token to be ... but got ... instead
	CODE_FILE_INTEGRITY_FILE_INVALID_MUST_BE_CSC_FILE                            = "EC_119" // Error | File that was checked was not a CSC file or did not end in .csc
	CODE_FILE_INTEGRITY_FILE_INVALID_MUST_NOT_BE_DIRECTORY_DIR_UNSUPPORTED       = "EC_120" // Error | File was not a file, it was a directory, unsupported
	CODE_FILE_INTEGRITY_FILE_INVALID_MUST_HAVE_CODE_OR_LOGIC_INSIDE_FILE_NULL    = "EC_121" // Error | File when scanned and checked came back empty or did not load any external or internal symbols from other imports
	CODE_FILE_INTEGRITY_FILE_INVALID_FAILED_TO_IMPORT_OR_OPEN_FILE               = "EC_122" // Error | File failed to open
	CODE_FILE_INTEGRITY_FILE_INVALID_FAILED_TO_STAT                              = "EC_123" // Error | File for some reason could not grab statistics due to ..,
	CODE_FILE_INTEGRITY_FILE_INVALID_FILE_NAME_WAS_EMPTY_OR_NULL_CHEC_INPUT      = "EC_125" // Error | File that was ran through the SkyLine interpreter or loaded into REPL may have been null or empty
	CODE_FILE_INTEGRITY_FILE_INVALID_FILE_DOES_NOT_EXIST                         = "EC_126" // Error | File DOES NOT exist
	CODE_FILE_FAILED_INJECTION_FILE_FAILED_TO_LINK_DUE_TO_NULLERR                = "EC_124" // Error | Linker failed to inject code loaded from internal imports to external symbols due to ...
	CODE_FILE_MUST_HAVE_NEW_FUNCTION_AND_METHOD_CALLED_DEVELOPER_ERROR_IN_SYMBOL = "EC_125" // Error | Program failed due to FileCurrent.New() method not being called to assign a new file, this means a file was not set for parsing or compile time
)

var ErrorSymBolMap = map[string]func(Arguments ...string) string{
	CODE_PARSE_FLOAT_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type float ", Arguments[0])
	}, // Error | Could not parse float value
	CODE_PARSE_INT_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type integer ", Arguments[0])
	}, // Error | Could not parse integer value
	CODE_PARSE_STRING_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type string", Arguments[0])
	}, // Error | Could not parse string value
	CODE_PARSE_BOOL_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type boolean", Arguments[0])
	}, //
	CODE_PARSE_NULL_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type null", Arguments[0])
	}, //
	CODE_PARSE_ARRAY_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type array", Arguments[0])
	}, //
	CODE_PARSE_HASH_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type hash", Arguments[0])
	}, //
	CODE_PARSE_HASHKEY_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not make  (%s) a useable hash key ", Arguments[0])
	}, //
	CODE_PARSE_TYPE_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Data Type mismatch in concentration, variable argument, call argument or function argument (Mismatch of type: %s and %s) ", Arguments[0], Arguments[1])
	}, //
	CODE_PARSE_OPERATOR_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Invalid operator (%s) see Skyline__('OPERATOR') for more info", Arguments[0])
	}, //
	CODE_PARSE_IDENTIFIER_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("No parse function found for identifier (%s) ", Arguments[0])
	}, //
	CODE_PARSE_FUNCTION_ARGUMENTS_NOT_ENOUGH_ERROR:                               func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_PARSE_MACRO_INVALID_ERROR:                                               func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_PARSE_INDEX_OPERATOR_UNSUPPORTED_WITHIN_KEY_NOTE_ERROR:                  func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_PARSE_AST_MODIFICATION_TO_MACRO_UNSUPPORTED_METHOD_ERROR:                func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_NO_FUNCTIONS_OR_SYMBOLS_LOADED:                                          func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_WRONG_NUMBER_OF_ARGUMENTS:                                               func(Arguments ...string) string { return fmt.Sprintf("") }, ////////////////////
	CODE_PREFIX_PARSE_FUNCTION_INVALID_OR_UNFOUND_WITHIN_PARSER_AND_INTERPRETRR:  func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_EXPECT_PEEK_ERROR_DURING_CALL_TO_PEEK:                                   func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_FILE_INTEGRITY_FILE_INVALID_MUST_BE_CSC_FILE:                            func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_FILE_INTEGRITY_FILE_INVALID_MUST_NOT_BE_DIRECTORY_DIR_UNSUPPORTED:       func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_FILE_INTEGRITY_FILE_INVALID_MUST_HAVE_CODE_OR_LOGIC_INSIDE_FILE_NULL:    func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_FILE_INTEGRITY_FILE_INVALID_FAILED_TO_IMPORT_OR_OPEN_FILE:               func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_FILE_INTEGRITY_FILE_INVALID_FAILED_TO_STAT:                              func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_FILE_INTEGRITY_FILE_INVALID_FILE_NAME_WAS_EMPTY_OR_NULL_CHEC_INPUT:      func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_FILE_INTEGRITY_FILE_INVALID_FILE_DOES_NOT_EXIST:                         func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_FILE_FAILED_INJECTION_FILE_FAILED_TO_LINK_DUE_TO_NULLERR:                func(Arguments ...string) string { return fmt.Sprintf("") }, //
	CODE_FILE_MUST_HAVE_NEW_FUNCTION_AND_METHOD_CALLED_DEVELOPER_ERROR_IN_SYMBOL: func(Arguments ...string) string { return fmt.Sprintf("") }, //
}

// Color list by OS

// ______LINUX OPERATING SYSTEMS__________

const (
	ERROR_RED = "\033[38;5;160m"
	ERROR_MSG = "\033[38;5;123m"
)
