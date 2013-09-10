Client.Router.map () ->
  
  @resource 'build_edit'
  @resource 'build_edit', path: '/build_edit/:build_edit_id'
  @resource 'build_edit.edit', path: '/build_edit/:build_edit_id/edit'
  
  @resource 'builds'
  @resource 'build', path: '/build/:build_id'
  @resource 'build.edit', path: '/build/:build_id/edit'