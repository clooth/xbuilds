Client.BuildsRoute = Ember.Route.extend
  model: () ->
    @store.push('build', {
      name: "Foobar build",
      createdAt: new Date(),
      updatedAt: new Date()
    })

    @store.find('build')