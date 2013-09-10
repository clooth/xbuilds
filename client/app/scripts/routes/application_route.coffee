Client.ApplicationRoute = Ember.Route.extend
    # admittedly, this should be in IndexRoute
    model: -> ['red', 'yellow', 'blue']
