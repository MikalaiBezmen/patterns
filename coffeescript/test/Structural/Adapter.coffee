#//TestTeaBagAdaptation.java - testing the adapter

define ['Structural/Adapter/TeaBag',
'Structural/Adapter/TeaCup',
'Structural/Adapter/TeaBall',
'Structural/Adapter/LooseLeafTea'
],(
TeaBag,
TeaCup,
TeaBall,
LooseLeafTea
) ->

  'use strict'

  describe 'Test Structural Adapter', ->
    it 'Steeping tea bag', ->
      teaCup = new TeaCup()
      teaBag = new TeaBag()
      expect(teaCup.steepTeaBag(teaBag)).toEqual("tea bag is steeping in cup")
        
    it 'Steeping loose leaf tea', ->
      teaCup = new TeaCup()
      looseLeafTea = new LooseLeafTea()
      teaBall = new TeaBall(looseLeafTea)
      expect( teaCup.steepTeaBag(teaBall) ).toBeTruthy()
