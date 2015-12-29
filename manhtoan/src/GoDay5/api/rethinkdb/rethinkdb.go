package rethinkdb

import (
	r "github.com/dancannon/gorethink"
)

type Instance struct {
	opts    r.ConnectOpts
	session *r.Session

	db string
}

func NewInstance(opts r.ConnectOpts) (ins *Instance, err error) {
	ins = &Instance{
		opts: opts,
		db:   opts.Database,
	}

	ins.session, err = r.Connect(opts)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (this *Instance) Connect() (*r.Session, error) {
	session, err := r.Connect(this.opts)
	return session, err
}

func (this *Instance) DB() r.Term {
	return r.DB(this.db)
}

func (this *Instance) Exec(term r.Term) error {
	session, err := this.Connect()
	if err != nil {
		return err
	}
	return term.Exec(session)
}

func (this *Instance) Table(name string) r.Term {
	return r.DB(this.db).Table(name)
}

func (this *Instance) TableCreate(name string) r.Term {
	return r.DB(this.db).TableCreate(name)
}

func (this *Instance) TableDrop(name string) r.Term {
	return r.DB(this.db).TableDrop(name)
}

func (this *Instance) Run(term r.Term) (*r.Cursor, error) {
	session, err := this.Connect()
	if err != nil {
		return nil, err
	}

	return term.Run(session)
}

func (this *Instance) RunWrite(term r.Term) (r.WriteResponse, error) {
	session, err := this.Connect()
	if err != nil {
		return r.WriteResponse{}, err
	}
	return term.RunWrite(session)
}
