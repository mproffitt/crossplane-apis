import common

items = [
  {
    **common.dxr
    **{
        status = {
            hello: get(oxr, "spec.message")
          }
      }
  }
]
