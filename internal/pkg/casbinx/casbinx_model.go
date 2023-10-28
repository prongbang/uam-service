package casbinx

const (
	/*
		https://casbin.org/docs/supported-models
		p, alice, data1, read
		p, bob, data2, write
		p, data2_admin, data2, read
		p, data2_admin, data2, write
		g, alice, data2_admin

		or

		allowed := e.Enforcer(alice, data2, read)  // true
		allowed := e.Enforcer(alice, data2, write) // true
	*/
	ModelRbacPolicy = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`
	ModelRestPolicy = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = (g(r.sub, p.sub) || r.sub == p.sub) && (keyMatch(r.obj, p.obj) || keyMatch2(r.obj, p.obj)) && (r.act == p.act || regexMatch(r.act, p.act))
`

	ModelGrpcPolicy = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = (g(r.sub, p.sub) || r.sub == p.sub) && (keyMatch(r.obj, p.obj) || keyMatch2(r.obj, p.obj)) && r.act == p.act
`
)
