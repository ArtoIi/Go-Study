package reflex

import "reflect"

// func percorre(x interface{}, fn func(entrada string)) {
// 	valor := reflect.ValueOf(x) // ValorDe
// 	campo := valor.Field(0)     // Campo
// 	fn(campo.String())
// }
//-----------------------------------------
// func percorre(x interface{}, fn func(entrada string)) {
// 	valor := reflect.ValueOf(x)

// 	for i := 0; i < valor.NumField(); i++ {
// 		campo := valor.Field(i)
// 		fn(campo.String())
// 	}
// }
//-----------------------------------------
// func percorre(x interface{}, fn func(entrada string)) {
// 	valor := reflect.ValueOf(x)

// 	for i := 0; i < valor.NumField(); i++ {
// 		campo := valor.Field(i)

// 		if campo.Kind() == reflect.String { // Tipo
// 			fn(campo.String())
// 		}
// 	}
// }
//-----------------------------------------
// func percorre(x interface{}, fn func(entrada string)) {
// 	valor := reflect.ValueOf(x)

// 	for i := 0; i < valor.NumField(); i++ {
// 		campo := valor.Field(i)

// 		if campo.Kind() == reflect.String {
// 			fn(campo.String())
// 		}

// 		if campo.Kind() == reflect.Struct {
// 			percorre(campo.Interface(), fn)
// 		}
// 	}
// }
//----------------------------------------- REFATORADO ^^
// func percorre(x interface{}, fn func(entrada string)) {
//     valor := reflect.ValueOf(x)

//     for i := 0; i < valor.NumField(); i++ {
//         campo := valor.Field(i)

//         switch campo.Kind() {
//         case reflect.String:
//             fn(campo.String())
//         case reflect.Struct:
//             percorre(campo.Interface(), fn)
//         }
//     }
// }
//-------------------------
// func percorre(x interface{}, fn func(entrada string)) {
// 	valor := obtemValor(x)
// 	for i := 0; i < valor.NumField(); i++ {
// 		campo := valor.Field(i)

// 		switch campo.Kind() {
// 		case reflect.String:
// 			fn(campo.String())
// 		case reflect.Struct:
// 			percorre(campo.Interface(), fn)
// 		}
// 	}
// }

// func obtemValor(x interface{}) (valor reflect.Value) {
// 	valor = reflect.ValueOf(x)

// 	if valor.Kind() == reflect.Ptr {
// 		valor = valor.Elem()
// 	}
// 	return valor
// }
//-----------------------------
func percorre(x interface{}, fn func(entrada string)) {
	valor := obtemValor(x)

	quantidadeDeValores := 0
	var obtemCampo func(int) reflect.Value

	switch valor.Kind() {
	case reflect.String:
		fn(valor.String())
	case reflect.Struct:
		quantidadeDeValores = valor.NumField()
		obtemCampo = valor.Field
	case reflect.Slice:
		quantidadeDeValores = valor.Len()
		obtemCampo = valor.Index
	}

	for i := 0; i < quantidadeDeValores; i++ {
		percorre(obtemCampo(i).Interface(), fn)
	}
}

func obtemValor(x interface{}) reflect.Value {
	valor := reflect.ValueOf(x)

	if valor.Kind() == reflect.Ptr {
		valor = valor.Elem()
	}

	return valor
}
