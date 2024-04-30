package container

func (c *Container) Configure() error {
	for _, constructor := range constructors {
		if err := c.container.Provide(constructor); err != nil {
			return err
		}
	}

	return nil
}
