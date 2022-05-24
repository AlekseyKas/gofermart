package storage_test

// func TestDatabase_InitDB(t *testing.T) {
// 	type fields struct {
// 		con   *pgxpool.Pool
// 		loger logrus.FieldLogger
// 		DBURL string
// 		ctx   context.Context
// 	}
// 	type args struct {
// 		ctx   context.Context
// 		DBURL string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := &Database{
// 				con:   tt.fields.con,
// 				loger: tt.fields.loger,
// 				DBURL: tt.fields.DBURL,
// 				ctx:   tt.fields.ctx,
// 			}
// 			if err := d.InitDB(tt.args.ctx, tt.args.DBURL); (err != nil) != tt.wantErr {
// 				t.Errorf("Database.InitDB() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
