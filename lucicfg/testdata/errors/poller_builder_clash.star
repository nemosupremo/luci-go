luci.project(
    name = 'proj',
    buildbucket = 'cr-buildbucket.appspot.com',
    swarming = 'chromium-swarm.appspot.com',
    scheduler = 'luci-scheduler.appspot.com',
)
luci.recipe(
    name = 'noop',
    cipd_package = 'noop',
)
luci.bucket(name = 'b')
luci.builder(
    name = 'clashing name',
    bucket = 'b',
    recipe = 'noop',
)
luci.gitiles_poller(
    name = 'clashing name',
    bucket = 'b',
    repo = 'https://noop.com',
)

# Expect errors like:
#
# Traceback (most recent call last):
#   //testdata/errors/poller_builder_clash.star:17: in <toplevel>
#   ...
# Error: luci.triggerer("b/clashing name") is redeclared, previous declaration:
# Traceback (most recent call last):
#   //testdata/errors/poller_builder_clash.star:12: in <toplevel>
#   ...
