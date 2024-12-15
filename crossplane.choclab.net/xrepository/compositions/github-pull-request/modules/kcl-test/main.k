import common

items = [
  {
    **common.dxr
    **{
        status = {
            hello: common.get(common.oxr, "spec.message")
          }
      }
  }
]
