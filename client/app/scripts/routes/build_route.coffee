Client.BuildRoute = Ember.Route.extend
  model: (model) ->
    @store.find('build', model.build_id)
