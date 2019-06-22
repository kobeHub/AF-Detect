import sys
import os
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '3'
import tensorflow as tf

from cardio import EcgDataset
#from cardio.dataset import B, V, F
from cardio.models.dirichlet_model import DirichletModel, concatenate_ecg_batch
from cardio.pipelines import dirichlet_predict_pipeline


signal_path = sys.argv[1]

BASE_DIR = os.path.abspath(os.path.dirname(os.getcwd()))
MODEL_PATH = 'model'
BATCH_SIZE = 100

gpu_options = tf.GPUOptions(allow_growth=True)

template_predict_ppl = dirichlet_predict_pipeline(model_path=MODEL_PATH,
        batch_size=BATCH_SIZE,
        gpu_options=gpu_options
        )


predict_eds = EcgDataset(path=signal_path)
predict_ppl = (predict_eds >> template_predict_ppl).run()

print(str(predict_ppl.get_variable("predictions_list")[0]).replace("'", "\""))
