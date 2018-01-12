package db

func IsExist(key string) (bool, error) {
	if dEtcd == nil {
		if err := InitDB(); err != nil {
			return false, err
		}
	}

	values, err := Get(key)
	if err != nil {
		return false, err
	}

	if values == nil || len(values) == 0 {
		return false, nil
	}

	return true, nil
}
