Client.BuildEditRoute = Ember.Route.extend
  model: (model) ->
    @get('store').find('build', model.build_id);

