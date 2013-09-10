Client.Build = DS.Model.extend
    objectId:   DS.attr('string')
    name:       DS.attr('string')
    createdAt:  DS.attr('date')
    updatedAt:  DS.attr('date')

Client.Build.FIXTURES = [
  {
    id: 0,
    objectId: 'foo1',
    name: 'foo',
    created_at: "2013-09-10T15:35:52.025+03:00",
    updated_at: "2013-09-10T15:35:52.025+03:00"
  },
  
  {
    id: 1,
    objectId: 'foo2',
    name: 'foo',
    created_at: "2013-09-10T15:35:52.025+03:00",
    updated_at: "2013-09-10T15:35:52.025+03:00"
  }
  
]
