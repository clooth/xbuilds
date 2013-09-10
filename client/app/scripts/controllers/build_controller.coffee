Client.BuildController = Ember.ObjectController.extend
    model: (params) ->
        @store.find('build', params.build_id)

    save: () ->
        @store.commit()
        @get('target.router').transitionTo('builds.index')