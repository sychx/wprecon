package text

func IndexAny(entity interface{}, field, value string) int {
    var _, index = any(entity, field, value)

    return index
}