package gofb

func IsEqualResponse(a, b *Response) bool {
	if len(a.Data) != len(b.Data) {
		return false
	}
	for i := range a.Data {
		for j := range a.Data[i] {
			if a.Data[i][j] != b.Data[i][j] {
				return false
			}
		}
	}

	if a.Error != b.Error {
		return false
	}

	// if len(a.Error) != len(b.Error) {
	// 	return false
	// }

	// for i := range a.Error {
	// 	if a.Error[i] != b.Error[i] {
	// 		return false
	// 	}
	// }

	if len(a.Schema.Fields) != len(b.Schema.Fields) {
		return false
	}

	for i := range a.Schema.Fields {
		if a.Schema.Fields[i].Name != b.Schema.Fields[i].Name {
			return false
		}
	}

	return true
}
