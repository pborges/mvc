package common

// This file should probably be burned in a fire

//
//func ParseFormAndRunActions(r *http.Request, dst Prefixer) (errors []error) {
//	errors = make([]error, 0)
//	err := r.ParseForm()
//	if err != nil {
//		errors = append(errors, err)
//		return
//	}
//
//	typeOf := reflect.TypeOf(dst)
//	if typeOf.Kind() == reflect.Ptr {
//		typeOf = typeOf.Elem()
//	}
//
//	valueOf := reflect.ValueOf(dst)
//	if valueOf.Kind() == reflect.Ptr {
//		valueOf = valueOf.Elem()
//	}
//
//	// parse the fields
//	for i := 0; i < typeOf.NumField(); i++ {
//		n := typeOf.Field(i).Name
//		if isPublic(n) {
//			f := valueOf.FieldByName(n)
//			errors = append(errors, convertAssign(dst.GetPrefix() + "." + n, r, f)...)
//		}
//	}
//
//	// parse the actions
//	args := strings.Split(r.FormValue(dst.GetPrefix() + "._action"), ".")
//	action, ok := reflect.TypeOf(dst).MethodByName(args[0])
//	if ok {
//		reflect.ValueOf(dst).MethodByName(action.Name).Call(nil)
//	}
//	return
//}
//
//func convertAssign(name string, r *http.Request, dst reflect.Value) (errors []error) {
//	errors = make([]error, 0)
//	switch reflect.TypeOf(dst.Interface()).Kind(){
//	case reflect.String:
//		value := r.FormValue(name)
//		dst.SetString(value)
//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//		value := r.FormValue(name)
//		i, err := strconv.ParseInt(value, 10, 64)
//		if err != nil {
//			errors = append(errors, err)
//		}else {
//			dst.SetInt(i)
//		}
//	//case reflect.Slice:
//	//	s := reflect.ValueOf(dst.Interface()).CallSlice()
//	//	notes := r.Form[name + "[]"]
//	//	this.Notes = make([]string, len(notes))
//	//	for i, note := range notes {
//	//		this.Notes[i] = note
//	//	}
//	default:
//		fmt.Errorf("Unknown type %s, %s\n", name, reflect.TypeOf(dst.Interface()).Kind())
//	}
//	return
//}
//
//func isPublic(name string) (bool) {
//	return strings.ToUpper(name[:1]) == name[:1]
//}