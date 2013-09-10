Client.BuildRoute = Ember.Route.extend
  model: (params) ->
    @store.find('build', params.build_id)
