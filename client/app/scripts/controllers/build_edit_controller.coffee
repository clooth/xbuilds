Client.BuildEditController = Ember.ObjectController.extend
  save: () ->
    @transitionToRoute 'build', @get('model')

