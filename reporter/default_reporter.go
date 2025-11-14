package reporter

type Default struct{}

func (d *Default) ReportMiss() {}

func (d *Default) ReportHit() {}
