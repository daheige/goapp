# validator-gen

    pb/validator.go
    package pb
    
    import (
    	"github.com/go-playground/validator/v10"
    )
    
    var validate = validator.New()
    
    pb/validator_req.go
    package pb
    
    // Validate req validator.
    func (r *HelloReq) Validate() error {
    	return validate.Struct(r)
    }
