Client.BuildsRoute = Ember.Route.extend
  model: () ->
    @store.find('build')
