package impl

import (
	"context"
	"go_learn/restful-api-demo/app/host"
)

func (i *impl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	if err := ins.Validate(); err != nil {
		return nil, err
	}
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	resStmt, err := tx.Prepare(insertResourceSQL)
	if err != nil {
		return nil, err
	}
	defer resStmt.Close()
	if _, err = resStmt.Exec(
		ins.Id, ins.Vendor, ins.Region, ins.Zone, ins.CreateAt, ins.ExpireAt, ins.Category, ins.Type, ins.InstanceId,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.SyncAccount, ins.PublicIP,
		ins.PrivateIP, ins.PayType, ins.ResourceHash, ins.DescribeHash,
	); err != nil {
		return nil, err
	}

	desStmt, err := tx.Prepare(insertDescribeSQL)
	defer desStmt.Close()
	if err != nil {
		return nil, err
	}
	if _, err = desStmt.Exec(ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec, ins.OSType, ins.OSName,
		ins.SerialNumber, ins.ImageID, ins.InternetMaxBandwidthOut,
		ins.InternetMaxBandwidthIn, ins.KeyPairName, ins.SecurityGroups); err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			err := tx.Rollback()
			i.log.Debugf("tx rollback err, %s", err)
		} else {
			err := tx.Commit()
			if err != nil {
				i.log.Debugf("tx commit err, %s", err)
			}
		}
	}()

	return nil, nil

}

func (i *impl) QUeryHost(ctx context.Context, req *host.QueryHostRequest) (*host.Set, error) {
	return nil, nil

}

func (i *impl) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (*host.Host, error) {
	return nil, nil

}

func (i *impl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil

}

func (i *impl) DeleteHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil

}
