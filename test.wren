
#!error = "Throws an error"
class Error {
  message { _message }
  construct new() {
    _message = "error"
  }
  construct new(msg) {
    _message = msg
  }
  #!doc="Returns error message"
  Error() {
    return message
  }
  toString {
    return message
  }
}

class FileNotFoundError is Error {
  construct new(msg) {
    super(msg)
  }
}


var err = Error.new()
var cls = err.type
System.print("docstring: %(cls.attributes.methods["Error()"][null]["doc"].toString())")