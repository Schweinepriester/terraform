variable "a" {
}

resource "test_thing" "foo" {
}

output "a" {
  value = {
    a   = var.a
    foo = test_thing.foo
  }
}
