package solutions

type AllOne struct {
    firstMetadata  map[string]int
    secondMetadata map[int]map[string]interface{}
    maxValue int
    minValue int
}

func Constructor() AllOne {
    return AllOne{
        firstMetadata:  map[string]int{},
        secondMetadata: map[int]map[string]interface{}{},
        maxValue:       0,
        minValue:       0,
    }
}

func (ao *AllOne) Inc(key string)  {
    ao.firstMetadata[key]++

    if ao.secondMetadata[ao.firstMetadata[key]] == nil {
        ao.secondMetadata[ao.firstMetadata[key]] = map[string]interface{}{}
    }

    ao.secondMetadata[ao.firstMetadata[key]][key] = nil

    delete(ao.secondMetadata[ao.firstMetadata[key] - 1], key)

    if ao.firstMetadata[key] > ao.maxValue {
        ao.maxValue++
    }

    if len(ao.secondMetadata[ao.minValue]) == 0 {
        ao.minValue++
    }

    if len(ao.secondMetadata[1]) != 0 {
        ao.minValue = 1
    }
}

func (ao *AllOne) Dec(key string)  {
    ao.firstMetadata[key]--

    if ao.firstMetadata[key] < 0 {
        delete(ao.firstMetadata, key)

        return
    }

    if ao.firstMetadata[key] != 0 {
        ao.secondMetadata[ao.firstMetadata[key]][key] = nil
    } else {
        delete(ao.firstMetadata, key)
    }

    delete(ao.secondMetadata[ao.firstMetadata[key] + 1], key)

    if len(ao.secondMetadata[ao.maxValue]) == 0 {
        ao.maxValue--
    }

    if ao.minValue > 1 {
        ao.minValue--
    } else if len(ao.firstMetadata) == 0 {
        ao.minValue = 0
    } else if len(ao.secondMetadata[1]) == 0 {
        ao.minValue = ao.maxValue

        for _, value := range ao.firstMetadata {
            if value < ao.minValue {
                ao.minValue = value
            }
        }
    }
}

func (ao *AllOne) GetMaxKey() string {
    for key := range ao.secondMetadata[ao.maxValue] {
        return key
    }

    return ""
}

func (ao *AllOne) GetMinKey() string {
    for key := range ao.secondMetadata[ao.minValue] {
        return key
    }

    return ""
}
