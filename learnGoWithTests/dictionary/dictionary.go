package dictionary

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrWordExists       = DictionaryErr("word already exists")
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newdefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newdefinition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string){
		delete(d, word)
}
