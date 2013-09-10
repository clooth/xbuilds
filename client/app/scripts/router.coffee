Client.Router.map () ->
  # Builds Route Mapping
  @resource 'builds'
  @resource 'build',            path: '/build/:build_id'
  @resource 'build.edit',       path: '/build/:build_id/edit'
