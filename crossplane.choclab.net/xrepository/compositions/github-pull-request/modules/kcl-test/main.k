import common

items = [
  {
    **common.dxr
    **{
        status = {
            hello: get(common.oxr, "spec.message")
          }
      }
  }
]
