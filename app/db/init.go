package db

func (p *Provider) Init() error {
	all := p.AllCollections()
	for _, v := range all {
		initable, ok := v.(interface{ Init() error })
		if !ok {
			continue
		}
		err := initable.Init()
		if err != nil {
			return err
		}
	}
	return nil
}