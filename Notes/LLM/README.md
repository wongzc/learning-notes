LLM learning note from **LLMs-from-scratch** by sebastian Raschka
https://github.com/rasbt/LLMs-from-scratch


1. use GPU for pytorch
    - run `nvidia-smi` to check if system is NVIDIA GPU, and check CUDA version
    - go to `https://pytorch.org/get-started/locally/` select and install
    - use below code to check if torch with cuda installed correctely
```python
import torch
print(torch.cuda.is_available())  # True means GPU is usable
print(torch.cuda.get_device_name())
```
