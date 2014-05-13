
features:
	( go run features/steps.go & ) && sleep .5 && cucumber --tags ~@wip

wip:
	( go run features/steps.go & ) && sleep .5 && cucumber --tags @wip --wip

.PHONY: features wip

